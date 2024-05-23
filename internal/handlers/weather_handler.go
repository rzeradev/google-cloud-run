package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rzeradev/google-cloud-run/internal/services"
)

func GetWeather(c *fiber.Ctx) error {
	zipcode := c.Params("zipcode")

	if len(zipcode) != 8 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "invalid zipcode",
		})
	}

	location, err := services.FetchLocation(zipcode)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "can not find zipcode",
		})
	}

	weather, err := services.FetchWeather(location.City)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to fetch weather data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(weather)
}
