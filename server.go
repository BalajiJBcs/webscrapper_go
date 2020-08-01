package main

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber"
	"log"
	"net/http"
	"context"
)

const dbName = "productdb"
const collectionName = "product"
const port = 8000

func getProduct(c *fiber.Ctx) {
	collection, err := getMongoDbCollection(dbName, collectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	var product Product

	json.Unmarshal([]byte(c.Body()), &product)
	client := &http.Client{}
	req, err := http.NewRequest("GET", product.WebUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var price string
	var review string

	name, _ := doc.Find("meta[name='title']").Attr("content")
	description, _ := doc.Find("meta[name='description']").Attr("content")
	price = doc.Find("#priceblock_ourprice").Text()
	review = doc.Find("#acrCustomerReviewText").Text()
	image, _ := doc.Find("#landingImage").Attr("data-old-hires")

	productFinal := ProductInfo{
		Name: name,
		Description: description,
		Price: price,
		Image: image,
		TotalReview: review,
	}
	res, err := collection.InsertOne(context.Background(), productFinal)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	response, _ := json.Marshal(res)
	c.Send(response)
}

func main() {
	app := fiber.New()
	app.Post("/product/", getProduct)
	app.Listen(port)
}

