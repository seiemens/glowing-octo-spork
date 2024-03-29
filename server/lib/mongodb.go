package lib

import (
	"context"
	"fmt"
	"fridge/models"
	"log"
	"time"

	"github.com/pquerna/otp/totp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func CreateUser(username, email, phone, password string, isAdmin bool) interface{} {
	var x = models.User{
		ID:       GenerateRandomString(8, false),
		Username: username,
		Email:    email,
		Phone:    phone,
		Password: HashAndSalt([]byte(password)),
		Admin:    false,
		ApiKey:   GenerateRandomString(16, false),
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

func GetSMSByKey(key string, value interface{}) []models.Smstoken {
	dbCollection := *Client.Database("fridge").Collection("smstokens")
	res, err := dbCollection.Find(context.Background(), bson.D{{key, primitive.Regex{Pattern: "^.*" + value.(string) + ".*", Options: ""}}})
	if err != nil {
		log.Fatal(err)
	}
	var result []models.Smstoken
	if err = res.All(context.Background(), &result); err != nil {
		log.Fatal(err)
	}
	return result
}

func GetSessionByKey(key string, value interface{}) []models.Cookie {
	dbCollection := *Client.Database("fridge").Collection("sessions")
	res, err := dbCollection.Find(context.Background(), bson.D{{key, primitive.Regex{Pattern: "^.*" + value.(string) + ".*", Options: ""}}})
	if err != nil {
		log.Fatal(err)
	}
	var result []models.Cookie
	if err = res.All(context.Background(), &result); err != nil {
		log.Fatal(err)
	}
	return result
}

func GetUserByKey(key string, value interface{}) []models.User {
	dbCollection := *Client.Database("fridge").Collection("userDB")
	res, err := dbCollection.Find(context.Background(), bson.D{{key, value.(string)}})
	if err != nil {
		log.Fatal(err)
	}
	var result []models.User
	if err = res.All(context.Background(), &result); err != nil {
		log.Fatal(err)
	}
	return result
}

func GetNaughtyByKey(key string, value interface{}) []models.Naughty {
	dbCollection := *Client.Database("fridge").Collection("naughtylist")
	res, err := dbCollection.Find(context.Background(), bson.D{{key, value.(string)}})
	if err != nil {
		log.Fatal(err)
	}
	var result []models.Naughty
	if err = res.All(context.Background(), &result); err != nil {
		log.Fatal(err)
	}
	return result
}

func GetNoteByStatus(key string, value models.Status) []models.Note {
	dbCollection := *Client.Database("fridge").Collection("noteDB")
	res, err := dbCollection.Find(context.Background(), bson.D{{key, value}})
	if err != nil {
		log.Fatal(err)
	}
	var result []models.Note
	if err = res.All(context.Background(), &result); err != nil {
		log.Fatal(err)
	}
	return result
}

func GetNoteByKey(key string, value interface{}) []models.Note {
	dbCollection := *Client.Database("fridge").Collection("noteDB")
	res, err := dbCollection.Find(context.Background(), bson.D{{key, primitive.Regex{Pattern: "^.*" + value.(string) + ".*", Options: ""}}})
	if err != nil {
		log.Fatal(err)
	}
	var result []models.Note
	if err = res.All(context.Background(), &result); err != nil {
		log.Fatal(err)
	}
	return result
}

func GetAllNotes() []models.Note {
	noteCollection := *Client.Database("fridge").Collection("noteDB")
	cursor, err := noteCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var notes []models.Note
	if err = cursor.All(context.Background(), &notes); err != nil {
		log.Fatal(err)
	}
	return notes
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
	if !CheckHashPassword(dbUser.Password, []byte(password)) {
		return models.User{}
	}
	return dbUser
}

func CreateSMSToken(phone, userID string) string {
	token := EncodeToString(6)
	tokenDB := Client.Database("fridge")
	tokenCollection := tokenDB.Collection("smstokens")
	t := models.Smstoken{
		ProcessID:   GenerateRandomString(6, false),
		UserID:      userID,
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
	return t.ProcessID
}

func VerifySMSToken(id, token string) (bool, string) {
	test := GetSMSByKey("process_id", id)

	if len(test) > 0 && test[0].AccessToken == token {
		return true, test[0].UserID
	} else {
		return false, ""
	}
}

func CreateSessionToken(userId string) models.Cookie {
	token := GenerateRandomString(12, false)
	tokenDB := Client.Database("fridge")
	tokenCollection := tokenDB.Collection("sessions")
	t := models.Cookie{
		UserID:    userId,
		Cookie:    token,
		CreatedAt: time.Now().UTC(),
		ExpireOn:  time.Now().Add(time.Hour).Unix(),
	}

	model := mongo.IndexModel{
		Keys:    bson.M{"created_at": 1},
		Options: options.Index().SetExpireAfterSeconds(3600), // delete session after 1h
	}

	_, err := tokenCollection.Indexes().CreateOne(context.Background(), model)

	if err != nil {
		log.Fatal(err)
	}
	_, err = tokenCollection.InsertOne(context.Background(), t)
	return t
}

func VerifySessionToken(sessionToken string) (bool, string) {
	test := GetSessionByKey("cookie", sessionToken)
	if len(test) > 0 {
		if test[0].Cookie == sessionToken {
			return true, test[0].UserID
		} else {
			return false, ""
		}
	} else {
		return false, ""
	}
}

func CreateNote(userID, title, content string) interface{} {
	user := GetUserByKey("id", userID)[0]

	var x = models.Note{
		ID:       GenerateRandomString(8, false),
		Title:    title,
		Content:  content,
		UserID:   userID,
		Author:   user.Username,
		Status:   models.Hidden,
		Comments: []models.Comment{},
	}
	noteDB := Client.Database("fridge")
	noteCollection := noteDB.Collection("noteDB")
	_, err := noteCollection.InsertOne(context.Background(), x)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully created note")
	return x

}

func IsAdmin(userID string) bool {
	user := GetUserByKey("id", userID)[0]
	return user.Admin
}

func AddCommentToPost(cookie, postID, content, userID string) {
	user := GetUserByKey("id", userID)[0]
	dbCollection := *Client.Database("fridge").Collection("noteDB")
	var addComment = models.Comment{
		ID:      GenerateRandomString(6, false),
		Content: content,
		PostID:  postID,
		Author:  user.Username,
	}
	fmt.Println(addComment)
	filter := bson.M{"id": bson.M{"$eq": postID}}
	change := bson.M{"$push": bson.M{"comments": addComment}}
	_, err := dbCollection.UpdateOne(
		context.Background(),
		filter,
		change,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func ChangePhone(number, userid string) {
	dbCollection := *Client.Database("fridge").Collection("userDB")
	filter := bson.M{"id": bson.M{"$eq": userid}}
	update := bson.M{"$set": bson.M{"phone": number}}

	_, err := dbCollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func Logout(cookie string) {
	tokenDB := Client.Database("fridge")
	tokenCollection := tokenDB.Collection("sessions")
	_, err := tokenCollection.DeleteOne(context.Background(), bson.M{"cookie": cookie})
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteSMSToken(pid string) {
	tokenDB := Client.Database("fridge")
	tokenCollection := tokenDB.Collection("smstokens")
	_, err := tokenCollection.DeleteOne(context.Background(), bson.M{"process_id": pid})
	if err != nil {
		log.Fatal(err)
	}
}

func ChangeVisibility(postID string, status models.Status) {
	tokenDB := Client.Database("fridge")
	noteDB := tokenDB.Collection("noteDB")
	filter := bson.M{"id": bson.M{"$eq": postID}}
	update := bson.M{"$set": bson.M{"status": status}}

	_, err := noteDB.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func ValidateAPIKey(key string) bool {
	user := GetUserByKey("apikey", key)
	if len(user) > 0 {
		return true
	} else {
		return false
	}
}

func IsUserNaughty(username string) bool {
	naughty := GetNaughtyByKey("username", username)
	return len(naughty) > 0
}

func CreateNaughtyOne(username string) models.Naughty {
	tokenDB := Client.Database("fridge")
	naughtyCollection := tokenDB.Collection("naughtylist")
	t := models.Naughty{
		Username:  username,
		CreatedAt: time.Now().UTC(),
		ExpireOn:  time.Now().Add(time.Minute * 5).Unix(),
	}

	model := mongo.IndexModel{
		Keys:    bson.M{"created_at": 1},
		Options: options.Index().SetExpireAfterSeconds(300), // delete session after 1h
	}

	_, err := naughtyCollection.Indexes().CreateOne(context.Background(), model)

	if err != nil {
		log.Fatal(err)
	}
	_, err = naughtyCollection.InsertOne(context.Background(), t)
	return t

}

func TOTP(inputTotp, userId string) bool {
	user := GetUserByKey("id", userId)[0]
	totp, err := totp.GenerateCode(user.ApiKey, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	return totp == inputTotp
}

func IsInitialSetup() bool {
	user := GetUserByKey("username", "Manfred")
	return !(len(user) > 0)
}

func InsertInitialData() {
	CreateUser("Manfred", "test@gmail.com", "", "Manfred123.", false)
	CreateUser("Michael", "michihimmelsberger@gmail.com", "", "Admin123.", true)
}
