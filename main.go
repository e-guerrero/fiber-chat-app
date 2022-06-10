package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   "1421342",
		Key:     "ddf41848e8ac318d4282",
		Secret:  "0b759ff1f76667db8086",
		Cluster: "us2",
		Secure:  true,
	}

	app.Post("/api/messages", func(c *fiber.Ctx) error {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		pusherClient.Trigger("chat", "message", data)
		return c.JSON([]string{})
	})

	/*
		No idea why PORT 3000 is being used as the defualt
		for Fiber instead of 8000 or 8080.
		All i found is that it's a standard in Express JS and
		Fiber is kinda like the Go version of Express JS.
		I'll be using 8000 only because the tutorial instructed
		it.
	*/
	app.Listen(":8000")
}
