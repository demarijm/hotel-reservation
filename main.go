package main

import (
	"flag"
	"log"

	"github.com/demarijm/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	port := flag.String("port", ":5001", "The listen address of the user")
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	log.Fatal(app.Listen(*port))
}
