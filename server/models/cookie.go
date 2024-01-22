package models

import "time"

type Cookie struct {
	UserID    string    `json:"user_id" bson:"user_id"`
	Cookie    string    `json:"cookie" bson:"cookie"`
	ExpireOn  int64     `json:"expire_on" bson:"expire_on"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
