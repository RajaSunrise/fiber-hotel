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


func GetAllFacility(c *fiber.Ctx) error{
	var facility []entity.Facility
	if err := database.DB.Find(&facility).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mendapatkan data facility",
		})
	}
	return c.JSON(fiber.Map{
		"facility": facility,
	})
}

func GetFacilityByID(c *fiber.Ctx) error{
	id := c.Params("id")
	var facility []entity.Facility
	if err := database.DB.Where("id = ?", id).First(&facility).Error; err != nil{
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "facility tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mendapatkan facility",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"facility": facility,
	})
}


func CreateFacility(c *fiber.Ctx) error{
	reqFacility := new(request.Facility)
	if err := c.BodyParser(reqFacility); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(reqFacility); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal membuat facility",
		})
	}

	if err := database.DB.Create(reqFacility).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal membuat facility",
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "berhasil membuat facility",
		"facility": reqFacility, 
	})
}


func UpdateFacility(c *fiber.Ctx) error{
	id := c.Params("id")
	var facility entity.Facility
	if err := database.DB.Where("id = ?", id).First(&facility).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "gagal menemukan data facility",
		})
	}
	updateFacility := &request.Facility{}
	if err := c.BodyParser(updateFacility); err != nil{
		return err
	}

	validate := validator.New()
	if err := validate.Struct(updateFacility); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal memperbarui facility",
		})
	}

	if err := database.DB.Save(&updateFacility).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal memperbarui facility",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "berhasil memperbarui facility",
		"facility": updateFacility,
	})
}


func DeleteFacility(c *fiber.Ctx) error{
	id := c.Params("id")
	var facility entity.Facility
	if err := database.DB.Where("id = ?", id).First(&facility).Error; err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "facility tidak ditemukan",
		})
	}
	if err := database.DB.Delete(&facility).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal menghapus facility",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("berhasil menghapus facility ID %s", id),
	})

}