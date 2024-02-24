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



func GetAllContent(c *fiber.Ctx) error{
	var content []entity.Content
	if err := database.DB.Find(&content).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal menemukan data content",
		})
	}
	return c.JSON(fiber.Map{
		"content": content,
	})
}


func GetContentByID(c *fiber.Ctx) error{
	id := c.Params("id")
	var content []entity.Content
	if err := database.DB.Where("id = ?", id).First(&content).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "content tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mendapatkan content",
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"content": content,
	})

}

func CreateContent(c *fiber.Ctx) error{
	reqContent := new(request.Content)
	if err := c.BodyParser(reqContent); err != nil{
		return err
	}

	validate := validator.New()
	if err := validate.Struct(reqContent); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal membuat Content",
		})
	}
	if err := database.DB.Create(reqContent).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal membuat data",
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "berhasil membuat content",
		"Content": reqContent,
	})
}

func UpdateContent(c *fiber.Ctx) error{
	id := c.Params("id")
	var content []entity.Content
	if err := database.DB.Where("id = ?", id).First(&content).Error; err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "gagal menemukan data content",
		})
	}
	updateContent := &request.Content{}
	if err := c.BodyParser(updateContent); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(&updateContent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal memperbarui content",
		})
	}

	if err := database.DB.Save(&updateContent).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal memperbarui content",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "berhasil memperbarui content",
		"content": updateContent,
	})
}


func DeleteContent(c *fiber.Ctx) error{
	id := c.Params("id")
	var content entity.Content
	if err := database.DB.Where("id = ?", id).First(&content).Error; err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "content tidak ditemukan",
		})
	}

	if err := database.DB.Delete(&content).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"gagal mengahapus content",
			"error": err.Error,
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("berhasil menghapus content ID %s", id),
	})
}