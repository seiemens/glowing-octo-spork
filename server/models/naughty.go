package models

import "time"

type Naughty struct {
	Username  string    `json:"username"`
	ExpireOn  int64     `json:"expire_on" bson:"expire_on"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
