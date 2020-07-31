package main

import (
	"encoding/json"
	"github.com/gofiber/fiber"
	"log"
)

const port = 8000


func getProduct(c *fiber.Ctx) {
	log.Println(c.Params("id"))
	var product Product
	json.Unmarshal([]byte(c.Body()), &product)
	log.Println(product.WebUrl)
	log.Println(product.Type)
}

func main() {
 log.Println("App Logged")
	app := fiber.New()
	app.Post("/product/", getProduct)
	app.Listen(port)
}

