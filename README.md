### Web Scrapper in Go

Simple Web URL Scrapper using the Go language, MongoDB, Fiber, GoQuery and Docker

### Requirement
 - Go
 - Docker

### Initializing the App

```dockerfile
docker-compose up -d --build
```

Login to Mongo Shell
```dockerfile
docker exec -it mongo mongo
```

HTTP Verbs

POST URL
```javascript
POST => localhost:8080/product/
```
POST BODY

Body Type raw and JSON

```
{
    "WebUrl": "YOUR WEBSITE URL"
}
```

Example Created Doc

```
{
	"_id" : ObjectId("5f252c5aa23bf37678a755e6"),
	"name" : "Title Goes Here",
	"image" : "URL Image Goes Here",
	"description" : "Description goes here",
	"price" : "₹ price",
	"totalreview" : "2,025 ratings",
	"created_at" : ISODate("2020-08-01T09:05:48.957Z"),
	"updated_at" : ISODate("2020-08-01T09:05:48.957Z")
}
```

GET ALL DETAILS OF PRODUCT 

```
GET => localhost:8080/product/
```

GET PARITCULAR PRODUCT ID

```
GET => localhost:8080/product/<PRODUCT_ID/>
```