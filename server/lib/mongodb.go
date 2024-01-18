package lib

import (
	"context"
	"fmt"
	"fridge/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client
var Uri = "mongodb://localhost:27017"

func ConnectToDb() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Uri))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(err)
	}
	Client = client
	fmt.Println("Successfully connected and pinged.")
}

func CloseDbConnection(client *mongo.Client) {
	var err error
	if err = client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func CreateUser(username, email, phone, password string) interface{} {
	var x = models.User{
		ID:       GenerateRandomString(8, false),
		Username: username,
		Email:    email,
		Phone:    phone,
		Password: password,
		Admin:    false,
	}
	userDB := Client.Database("fridge")
	userCollection := userDB.Collection("userDB")
	_, err := userCollection.InsertOne(context.Background(), x)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully created user")
	return x

}

func AuthUser(username, password string) models.User {
	var dbUser models.User
	userCollection := Client.Database("fridge").Collection("userDB")
	filter := bson.D{{"username", username}}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&dbUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}
		}
		log.Fatal(err)
	}
	if !CheckHashPassword(dbUser.Password, password) {
		return models.User{}
	}
	return dbUser
}

func CreateSMSToken(phone string) {
	token := EncodeToString(6)
	tokenDB := Client.Database("fridge")
	tokenCollection := tokenDB.Collection("smstokens")
	t := models.Smstoken{
		AccessToken: token,
		CreatedAt:   time.Now().UTC(),
		ExpireOn:    time.Now().Add(time.Second * 10).Unix(),
	}

	model := mongo.IndexModel{
		Keys:    bson.M{"created_at": 1},
		Options: options.Index().SetExpireAfterSeconds(300), // delete totp after 5 min
	}

	_, err := tokenCollection.Indexes().CreateOne(context.Background(), model)

	if err != nil {
		log.Fatal(err)
	}
	_, err = tokenCollection.InsertOne(context.Background(), t)
	SendSMS(phone, token)
}
