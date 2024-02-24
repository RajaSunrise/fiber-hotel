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


func GetAllLocation(c *fiber.Ctx) error{
	var location []entity.Location
	if err := database.DB.Find(&location).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mendapatkan data location",
		})
	}
	return c.JSON(fiber.Map{
		"location": location,
	})
}

func GetLocationByID(c *fiber.Ctx) error{
	id := c.Params("id")
	var location []entity.Location
	if err := database.DB.Where("id = ?", id).First(&location).Error; err != nil{
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "location tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mendapatkan location",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"location": location,
	})
}


func CreateLocation(c *fiber.Ctx) error{
	reqLocation := new(request.Location)
	if err := c.BodyParser(reqLocation); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(reqLocation); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal membuat location",
		})
	}

	if err := database.DB.Create(reqLocation).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal membuat location",
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "berhasil membuat location",
		"location": reqLocation, 
	})
}


func UpdateLocation(c *fiber.Ctx) error{
	id := c.Params("id")
	var location entity.Location
	if err := database.DB.Where("id = ?", id).First(&location).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "gagal menemukan data location",
		})
	}
	updateLocation := &request.Location{}
	if err := c.BodyParser(updateLocation); err != nil{
		return err
	}

	validate := validator.New()
	if err := validate.Struct(updateLocation); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal memperbarui location",
		})
	}

	if err := database.DB.Save(&updateLocation).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal memperbarui location",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "berhasil memperbarui location",
		"location": updateLocation,
	})
}


func DeleteLocation(c *fiber.Ctx) error{
	id := c.Params("id")
	var location entity.Location
	if err := database.DB.Where("id = ?", id).First(&location).Error; err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "location tidak ditemukan",
		})
	}
	if err := database.DB.Delete(&location).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal menghapus location",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("berhasil menghapus location ID %s", id),
	})

}