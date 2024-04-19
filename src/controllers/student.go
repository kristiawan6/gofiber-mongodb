package controllers

import (
	models "gofiber-mongodb/src/models"

	"github.com/gofiber/fiber/v2"
)

func FindData(c *fiber.Ctx) error {
	student := models.Find()

	// Check if student data is empty
	if len(student) == 0 {
		// Return an error indicating no data found
		return fiber.NewError(fiber.StatusNotFound, "No student data found")
	}

	// Send the student data as response
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    student,
	})
}

func CreateData(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost {
		var Student models.Student
		if err := c.BodyParser(&Student); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		item := models.Student{
			Name:  Student.Name,
			Grade: Student.Grade,
		}

		models.Insert(&item)
		return c.JSON(fiber.Map{
			"Message": "success",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not allowed")
	}
}

func UpdateData(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPut {
		nameParam := c.Params("name")
		var Student models.Student
		if err := c.BodyParser(&Student); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		newStudent := models.Student{
			Name:  Student.Name,
			Grade: Student.Grade,
		}

		models.Update(nameParam,&newStudent)
		return c.JSON(fiber.Map{
			"Message": "success",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not allowed")
	}
}

func DeleteData(c *fiber.Ctx) error {
	nameParam := c.Params("name")
	models.Delete(nameParam)

	return c.JSON(fiber.Map{
		"Message": "Deleted Success",
	})
}
