package controllers

import (
	"context"
	"github.com/Kosodaka/auth-service/database"
	"github.com/Kosodaka/auth-service/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func GetTokens(c *fiber.Ctx) error {
	var err error
	var input map[string]string
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	access, err := models.GenerateAccessToken(input["guid"])
	if err != nil {
		log.Println("failed to create access token")
		return err
	}
	refresh, err := models.GenerateRefreshToken()
	if err != nil {
		log.Println("failed to create refresh token")
		return err
	}
	refreshHash, _ := bcrypt.GenerateFromPassword([]byte(refresh), 14)
	database.Collection.InsertOne(context.TODO(), models.User{Guid: input["guid"], Refresh: string(refreshHash), Time: time.Now().Add(30 * 24 * time.Hour)})

	return c.JSON(map[string]interface{}{
		"guid":    input["guid"],
		"access":  access,
		"refresh": refreshHash})
}

func UpdateTokens(c *fiber.Ctx) error {
	var input map[string]string
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	var newAccess string
	var newRefresh string
	var status int
	var err error
	err = database.ValidateRefreshToken(input["guid"], input["refresh"])
	if err == nil {
		newAccess, err = models.GenerateAccessToken(input["guid"])
		if err != nil {
			log.Println("failed to create na")
		}
		newRefresh, _ = models.GenerateRefreshToken()
	}

	refreshHash, _ := bcrypt.GenerateFromPassword([]byte(newRefresh), 14)
	err = database.UpdateTokens(input["guid"], string(refreshHash))
	if err != nil {
		status = 0
	} else {
		status = 1
	}

	return c.JSON(map[string]interface{}{
		"newAccess":  newAccess,
		"newRefresh": refreshHash,
		"status":     status},
	)
}
