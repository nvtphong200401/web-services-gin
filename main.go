package main

import (
	"example/web-services-gin/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", handlers.ListLandmarkHandler)
	router.Run()
}
