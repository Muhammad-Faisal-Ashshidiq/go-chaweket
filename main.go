package main

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/sidiq200/go-chaweket/module"
	"github.com/sidiq200/go-chaweket/url"
	"log"
)

func main() {
	go module.Run()

	app := fiber.New()
	url.Web(app)
	log.Fatal(app.Listen(musik.Dangdut()))
}
