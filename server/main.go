package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", index)

	v1 := router.Group("/bot")
	{
		v1.POST("/request", ApiBotRequest)
	}

	log.Fatal(router.Run(":8080"))
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Admin Page",
	})
}

func ApiBotRequest(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "bot",
	})
}
