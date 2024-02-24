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


func GetAllPayment(c *fiber.Ctx) error{
	var payment []entity.Payment
	if err := database.DB.Find(&payment).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mendapatkan data payment",
		})
	}
	return c.JSON(fiber.Map{
		"payment": payment,
	})
}

func GetPaymentByID(c *fiber.Ctx) error{
	id := c.Params("id")
	var payment []entity.Payment
	if err := database.DB.Where("id = ?", id).First(&payment).Error; err != nil{
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "payment tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mendapatkan payment",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"payment": payment,
	})
}


func CreatePayment(c *fiber.Ctx) error{
	reqPayment := new(request.Payment)
	if err := c.BodyParser(reqPayment); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(reqPayment); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal membuat payment",
		})
	}

	if err := database.DB.Create(reqPayment).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal membuat payment",
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "berhasil membuat payment",
		"payment": reqPayment, 
	})
}


func Updatepayment(c *fiber.Ctx) error{
	id := c.Params("id")
	var payment entity.Payment
	if err := database.DB.Where("id = ?", id).First(&payment).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "gagal menemukan data payment",
		})
	}
	updatePayment := &request.Payment{}
	if err := c.BodyParser(updatePayment); err != nil{
		return err
	}

	validate := validator.New()
	if err := validate.Struct(updatePayment); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal memperbarui payment",
		})
	}

	if err := database.DB.Save(&updatePayment).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal memperbarui payment",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "berhasil memperbarui payment",
		"payment": updatePayment,
	})
}


func DeletePayment(c *fiber.Ctx) error{
	id := c.Params("id")
	var payment entity.Payment
	if err := database.DB.Where("id = ?", id).First(&payment).Error; err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "payment tidak ditemukan",
		})
	}
	if err := database.DB.Delete(&payment).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal menghapus payment",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("berhasil menghapus payment ID %s", id),
	})

}