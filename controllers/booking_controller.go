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



func GetAllBooking(c *fiber.Ctx) error{
	var booking []entity.Booking
	if err := database.DB.Find(&booking).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal menemukan data facility",
		})
	}
	return c.JSON(fiber.Map{
		"booking": booking,
	})
}


func GetBookingByID(c *fiber.Ctx) error{
	id := c.Params("id")
	var booking []entity.Booking
	if err := database.DB.Where("id = ?", id).First(&booking).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "booking tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mendapatkan booking",
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"booking": booking,
	})

}

func CreateBooking(c *fiber.Ctx) error{
	reqBooking := new(request.Booking)
	if err := c.BodyParser(reqBooking); err != nil{
		return err
	}

	validate := validator.New()
	if err := validate.Struct(reqBooking); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal membuat booking",
		})
	}
	if err := database.DB.Create(reqBooking).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal membuat data",
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "berhasil membuat bookig",
		"booking": reqBooking,
	})
}

func UpdateBooking(c *fiber.Ctx) error{
	id := c.Params("id")
	var booking []entity.Booking
	if err := database.DB.Where("id = ?", id).First(&booking).Error; err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "gagal menemukan data booking",
		})
	}
	updateBooking := &request.Booking{}
	if err := c.BodyParser(updateBooking); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(&updateBooking); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal memperbarui booking",
		})
	}

	if err := database.DB.Save(&updateBooking).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal memperbarui booking",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "berhasil memperbarui booking",
		"booking": updateBooking,
	})
}


func DeleteBooking(c *fiber.Ctx) error{
	id := c.Params("id")
	var booking entity.Booking
	if err := database.DB.Where("id = ?", id).First(&booking).Error; err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "booking tidak ditemukan",
		})
	}

	if err := database.DB.Delete(&booking).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"gagal mengahapus booking",
			"error": err.Error,
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("berhasil menghapus booking ID %s", id),
	})
}