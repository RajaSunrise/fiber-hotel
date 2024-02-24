package controllers

import (
	"github.com/RajaSunrise/hotel/database"
	"github.com/RajaSunrise/hotel/models/request"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)


func Register(c *fiber.Ctx) error {
	reqUser := new(request.User)
	if err := c.BodyParser(reqUser); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(reqUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal membuat user",
			"error":   err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal membuat user",
			"error":   err.Error(),
		})
	}

	reqUser.Password = string(hashedPassword)

	if err := database.DB.Create(reqUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal membuat user",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "berhasil membuat user",
		"data":    reqUser,
	})
}