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


func GetAllRoom(c *fiber.Ctx) error{
	var room []entity.Room
	if err := database.DB.Find(&room).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mendapatkan daftar room",
		})
	}
	return c.JSON(fiber.Map{
		"room": room,
	})
}

func GetRoomByID(c *fiber.Ctx) error{
	id := c.Params("id")
	var room []entity.Room
	if err := database.DB.Where("id = ?", id).First(&room).Error; err != nil{
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "room tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mendapatkan room",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"room": room,
	})
}

func CreateRoom(c *fiber.Ctx) error{
	reqRoom := new(request.Room)
	if err := c.BodyParser(reqRoom); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(reqRoom); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal membuat room",
		})
	}

	if err := database.DB.Create(reqRoom).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal memebuat room",
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "berhasil membuat room",
		"room": reqRoom, 
	})
}

func UpdateRoom(c *fiber.Ctx) error{
	id := c.Params("id")
	var room entity.Room
	if err := database.DB.Where("id = ?", id).First(&room).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "gagal menemukan data room",
		})
	}
	updateRoom := &request.Room{}
	if err := c.BodyParser(updateRoom); err != nil{
		return err
	}

	validate := validator.New()
	if err := validate.Struct(updateRoom); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal memperbarui room",
		})
	}

	if err := database.DB.Save(&updateRoom).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal memperbarui room",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "berhasil memperbarui room",
		"room": updateRoom,
	})
}


func DeleteRoom(c *fiber.Ctx) error{
	id := c.Params("id")
	var room entity.Room
	if err := database.DB.Where("id = ?", id).First(&room).Error; err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "room tidak ditemukan",
		})
	}
	if err := database.DB.Delete(&room).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal menghapus room",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("berhasil menghapus room ID %s", id),
	})

}