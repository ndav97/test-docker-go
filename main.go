package main

import (
    "github.com/gofiber/fiber"

    "os"
)

func main() {
  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) {
    c.Send("Hollaaassss")
  })

  app.Listen(os.Getenv("PORT"))
}
