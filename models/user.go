package models

import (
	"time"
)

type User struct {
	Guid    string    `bson:"_id"`
	Refresh string    `bson:"refresh"`
	Time    time.Time `json:"time" bson:"time" `
}
