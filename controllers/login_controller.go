package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/RajaSunrise/hotel/database"
	"github.com/RajaSunrise/hotel/middleware"
	"github.com/RajaSunrise/hotel/models/entity"
	"github.com/RajaSunrise/hotel/models/request"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)



func Login(c *fiber.Ctx) error {
	var reqUser request.User
	if err := c.BodyParser(&reqUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	var user entity.User
	if err := database.DB.Where("username = ? OR email = ?", reqUser.Username, reqUser.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqUser.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong password",
		})
	}

	// Set claims
	claims := jwt.MapClaims{
		"username": reqUser.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Membuat token dengan metode HS256 dan claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Menandatangani token dengan secret key
	t, err := token.SignedString(middleware.JwtSecret)
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("berhasil masuk sebagai %s", user.Username),
		"token":   t,
	})
}


