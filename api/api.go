package api

import (
	"strconv"

	"github.com/Khaliiloo/pack-size/packsize"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Api() *fiber.App {
	app := fiber.New()
	corsConfig := cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}
	app.Use(cors.New(corsConfig))

	app.Get("/numberOfPacks/:packsize", func(c *fiber.Ctx) error {
		stringSize := c.Params("packsize")
		ps, err := strconv.Atoi(stringSize)
		if err != nil {
			c.Status(500)
			c.JSON("fail")
		}
		result := packsize.NumberOfPacks(ps)
		return c.JSON(result)
	})

	app.Post("/addPackSize/:packsize", func(c *fiber.Ctx) error {
		stringSize := c.Params("packsize")
		ps, err := strconv.Atoi(stringSize)
		if err != nil {
			c.Status(500)
			c.JSON("fail")
		}
		packsize.AddPackSize(ps)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Delete("/removePackSize/:sizepack", func(c *fiber.Ctx) error {
		stringSize := c.Params("packsize")
		ps, err := strconv.Atoi(stringSize)
		if err != nil {
			c.Status(500)
			c.JSON("fail")
		}
		packsize.RemovePackSize(ps)
		return c.SendStatus(fiber.StatusNoContent)
	})
	return app
}
