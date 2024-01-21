package models

import "time"

type Smstoken struct {
	ProcessID   string    `json:"process_id" bson:"process_id"`
	UserID      string    `json:"user_id" bson:"user_id"`
	AccessToken string    `json:"access_token" bson:"access_token"`
	ExpireOn    int64     `json:"expire_on" bson:"expire_on"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
