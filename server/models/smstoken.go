package models

import "time"

type Smstoken struct {
	AccessToken string    `json:"access_token" bson:"access_token"`
	ExpireOn    int64     `json:"expire_on" bson:"expire_on"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
