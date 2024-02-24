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


func GetAllReview(c *fiber.Ctx) error{
	var review []entity.Review
	if err := database.DB.Find(&review).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal menampilkan data review",
		})
	}
	return c.JSON(fiber.Map{
		"review": review,
	})
}

func GetReviewByID(c *fiber.Ctx) error{
	id := c.Params("id")
	var review []entity.Review
	if err := database.DB.Where("id = ?", id).First(&review).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "review tidak ditemukan",
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "gagal mendapatkan data review",
		"error": err.Error(),
	})
	}
	return c.JSON(fiber.Map{
		"review": review,
	})
}


func CreateReview(c *fiber.Ctx) error{
	reqReview := new(request.Review)
	if err := c.BodyParser(&reqReview); err != nil{
		return err
	}

	validate := validator.New()
	if err := validate.Struct(reqReview); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal membuat review",
			"error": err.Error(),
		})
	}
	if err := database.DB.Create(reqReview).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal membuat review",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"review": reqReview,
	})
}

func UpdateReview(c *fiber.Ctx) error {
	id := c.Params("id")
	var review entity.Review
	if err := database.DB.Where("id = ?", id).First(&review).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "gagal menemukan data review",
		})
	}

	updateReview := &request.Review{}
	if err := c.BodyParser(updateReview); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(updateReview); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal memperbarui review",
		})
	}

	if err := database.DB.Save(&review).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal memperbarui review",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "berhasil memperbarui review",
		"review":  review,
	})
}

func DeleteReview(c *fiber.Ctx) error{
	id := c.Params("id")
	var review entity.Review
	if err := database.DB.Where("id = ?", id).First(&review).Error; err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "review tidak ditemukan",
		})
	}
	if err := database.DB.Delete(&review).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal menghapus review",
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("berhasil menghapus review ID %s", id),
	})
}
