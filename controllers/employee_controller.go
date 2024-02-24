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



func GetAllEmployee(c *fiber.Ctx) error{
	var employee []entity.Employee
	if err := database.DB.Find(&employee).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal menemukan data employee",
		})
	}
	return c.JSON(fiber.Map{
		"employee": employee,
	})
}


func GetEmployeeByID(c *fiber.Ctx) error{
	id := c.Params("id")
	var employee []entity.Employee
	if err := database.DB.Where("id = ?", id).First(&employee).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "employee tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mendapatkan employee",
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"employee": employee,
	})

}

func CreateEmployee(c *fiber.Ctx) error{
	reqEmployee := new(request.Employee)
	if err := c.BodyParser(reqEmployee); err != nil{
		return err
	}

	validate := validator.New()
	if err := validate.Struct(reqEmployee); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal membuat employee",
		})
	}
	if err := database.DB.Create(reqEmployee).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal membuat data",
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "berhasil membuat employee",
		"employee": reqEmployee,
	})
}

func UpdateEmployee(c *fiber.Ctx) error{
	id := c.Params("id")
	var employee []entity.Employee
	if err := database.DB.Where("id = ?", id).First(&employee).Error; err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "gagal menemukan data employee",
		})
	}
	updateEmployee := &request.Employee{}
	if err := c.BodyParser(updateEmployee); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(&updateEmployee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal memperbarui employee",
		})
	}

	if err := database.DB.Save(&updateEmployee).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal memperbarui employee",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "berhasil memperbarui employee",
		"employee": updateEmployee,
	})
}


func DeleteEmployee(c *fiber.Ctx) error{
	id := c.Params("id")
	var employee entity.Employee
	if err := database.DB.Where("id = ?", id).First(&employee).Error; err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "employee tidak ditemukan",
		})
	}

	if err := database.DB.Delete(&employee).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"gagal mengahapus employee",
			"error": err.Error,
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("berhasil menghapus employee ID %s", id),
	})
}