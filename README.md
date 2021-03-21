# Gin Tutorial

Following [this tutorial](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin).

It's [GitHub](https://github.com/demo-apps/go-gin-app) page.

Another one to look at. <https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/>

## Curl Commands

Get the API version:

```
curl -H "Accept: application/json" http://localhost:8080/api/ver
```

Now needs keys:

```
curl -X POST --data "apiKey=123" -H "apiKey: 123" http://localhost:8080/api/ver -H "Accept: application/json"
```

## Submit and handle JSON

```
curl -X POST http://localhost:8080/robin -H "Accept: application/json" -H "Content-Type: application/json" --data '{"url":"test.com"}'
```

## Testing

To test everything:

```
go test -v
```
