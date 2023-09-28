package database

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/Kosodaka/auth-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func FindId(guid string) (*models.User, error) {
	filter := bson.D{{"_id", guid}}
	var result *models.User
	var err error
	err = Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == nil {
		return result, err
	}
	return nil, err
}

func UpdateTokens(guid, refresh string) error {
	filter := bson.D{{"_id", guid}}
	update := bson.D{
		{"$set", bson.D{
			{"refresh", refresh},
			{"time", time.Now().Add(30 * 24 * time.Hour)},
		}},
	}
	_, err := Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
func ValidateRefreshToken(guid, refresh string) error {
	var err error
	var dbRef *models.User
	var decodeRef []byte
	if dbRef, err = FindId(guid); err == nil {
		if decodeRef, err = base64.StdEncoding.DecodeString(refresh); err == nil {
			if err = bcrypt.CompareHashAndPassword([]byte(dbRef.Refresh), decodeRef); err != nil {
				return nil
			} else {
				fmt.Println(string(dbRef.Refresh))
				fmt.Println(string(decodeRef))
			}
		}
	}
	return err
}
