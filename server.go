package main

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber"
	"log"
	"net/http"
)

const port = 8000

func getProduct(c *fiber.Ctx) {
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

	log.Println(doc)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(doc.Find("meta[name='title']").Attr("content"))
	log.Println(doc.Find("meta[name='description']").Attr("content"))
	log.Println(doc.Find("#priceblock_dealprice").Text())
	log.Println(doc.Find("#priceblock_ourprice").Text())
	log.Println(doc.Find("#acrCustomerReviewText").Text())
	log.Println(doc.Find("#landingImage").Attr("data-old-hires"))
}

func main() {
	app := fiber.New()
	app.Post("/product/", getProduct)
	app.Listen(port)
}

