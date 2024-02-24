package controllers

import (
	"fmt"

	"github.com/RajaSunrise/hotel/database"
	"github.com/RajaSunrise/hotel/models/entity"
	"github.com/RajaSunrise/hotel/models/request"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []entity.User
	if err := database.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mendapatkan daftar pengguna",
		})
	}
	return c.JSON(fiber.Map{
		"users": users,
	})
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var user []entity.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "pengguna tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mendapatkan pengguna",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"user": user,
	})
}


func CreateUser(c *fiber.Ctx) error{
	reqUser := new(request.User)
	if err := c.BodyParser(reqUser); err != nil{
		return err
	}

	validate := validator.New()
	if err := validate.Struct(reqUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal membuat user",
			"error":   err.Error(),
		})
	}
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


func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user entity.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user tidak ditemukan",
		})
	}

	updateUser := &request.User{}
	if err := c.BodyParser(updateUser); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(updateUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal memperbarui user",
			"error":   err.Error(),
		})
	}

	if err := database.DB.Model(&user).Updates(updateUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal memperbarui user",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "berhasil memperbarui user",
		"data":    updateUser,
	})
}


func DeleteUser(c *fiber.Ctx) error{
	id := c.Params("id")
	var user entity.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "pengguna tidak ditemukan",
		})
	}

	if err := database.DB.Delete(&user).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal menghapus users",
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("berhasil menghapus users ID %s", id),
	})
}