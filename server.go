package main

import (
	"context"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber"
	"github.com/parnurzeal/gorequest"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

const dbName = "productdb"
const collectionName = "product"
const port = 8000

func postProduct(c *fiber.Ctx) {
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
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	}
	request := gorequest.New()
	resp, body, errs := request.Post("http://localhost:8000/productInsert").
		Send(productFinal).
		End()
	if errs != nil {
		log.Fatal(err)
	}
	c.Send(body)
}

func insertProduct(c *fiber.Ctx) {
	collection, err := getMongoDbCollection(dbName, collectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	var products ProductInfo
	json.Unmarshal([]byte(c.Body()), &products)
	res, err := collection.InsertOne(context.Background(), products)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	response, _ := json.Marshal(res)
	c.Send(response)

}

func getProduct(c *fiber.Ctx) {
	collection, err := getMongoDbCollection(dbName, collectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var filter bson.M = bson.M{}

	if c.Params("id") != "" {
		id := c.Params("id")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404)
		return
	}

	json, _ := json.Marshal(results)
	c.Send(json)
}

func main() {
	app := fiber.New()
	app.Post("/product/", postProduct)
	app.Post("/productInsert", insertProduct)
	app.Get("/product/:id?", getProduct)
	app.Listen(port)
}

