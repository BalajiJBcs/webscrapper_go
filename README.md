### Web Scrapper in Go

Simple Web Scrapper using the Go language, MongoDB, Fiber, GoQuery and Docker

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
