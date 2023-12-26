package controller

import (
	"github.com/KaisarOrange/smart-office/model"
	"github.com/gofiber/fiber/v2"
)

func SetReminder(c *fiber.Ctx) error {
	reminder := new(model.Reminder)

	err:= c.BodyParser(&reminder)
	if err!=nil{
		c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
	}

	return c.Status(200).JSON(fiber.Map{"data":reminder})
}