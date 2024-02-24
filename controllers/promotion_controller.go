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


func GetAllPromotion(c *fiber.Ctx) error{
	var promotion []entity.Promotion
	if err := database.DB.Find(&promotion).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mendapatkan data promotion",
		})
	}
	return c.JSON(fiber.Map{
		"promotion": promotion,
	})
}

func GetpromotionByID(c *fiber.Ctx) error{
	id := c.Params("id")
	var promotion []entity.Promotion
	if err := database.DB.Where("id = ?", id).First(&promotion).Error; err != nil{
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "promotion tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mendapatkan promotion",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"promotion": promotion,
	})
}


func CreatePromotion(c *fiber.Ctx) error{
	reqPromotion := new(request.Promotion)
	if err := c.BodyParser(reqPromotion); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(reqPromotion); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal membuat promotion",
		})
	}

	if err := database.DB.Create(reqPromotion).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal membuat promotion",
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "berhasil membuat promotion",
		"promotion": reqPromotion, 
	})
}


func UpdatePromotion(c *fiber.Ctx) error{
	id := c.Params("id")
	var promotion entity.Promotion
	if err := database.DB.Where("id = ?", id).First(&promotion).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "gagal menemukan data promotion",
		})
	}
	updatePromotion := &request.Promotion{}
	if err := c.BodyParser(updatePromotion); err != nil{
		return err
	}

	validate := validator.New()
	if err := validate.Struct(updatePromotion); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal memperbarui promotion",
		})
	}

	if err := database.DB.Save(&updatePromotion).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal memperbarui promotion",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "berhasil memperbarui promotion",
		"promotion": updatePromotion,
	})
}


func DeletePromotion(c *fiber.Ctx) error{
	id := c.Params("id")
	var promotion entity.Promotion
	if err := database.DB.Where("id = ?", id).First(&promotion).Error; err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "promotion tidak ditemukan",
		})
	}
	if err := database.DB.Delete(&promotion).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal menghapus promotion",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("berhasil menghapus promotion ID %s", id),
	})

}