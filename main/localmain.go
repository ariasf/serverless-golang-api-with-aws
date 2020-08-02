package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"serverless-golang-api-with-aws/router"
)

func main() {
	r := gin.Default()
	router.LoadRoutes(r)

	if err := r.Run(); err != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		log.Printf("error starting server %+v", err)
	}
}
