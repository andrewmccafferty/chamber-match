package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

type Musician struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	v1 := router.Group("/api/v1/musician")
	{
		v1.POST("/", createMusician)
		v1.GET("/", fetchAllMusicians)
		v1.GET("/:id", fetchSingleMusician)
		v1.PUT("/:id", updateMusician)
		v1.DELETE("/:id", deleteMusician)
	}
	router.Run()

	router.Run(":" + port)
}

func createMusician(c *gin.Context) {
	var musician Musician
	if err := c.ShouldBindJSON(&musician); err == nil {
		log.Print(musician)
		c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("Got data %+v", musician)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "missing data"})
		log.Printf("couldn't bind because of error %+v", err)
	}
}

func fetchAllMusicians(c *gin.Context) {

}

func fetchSingleMusician(c *gin.Context) {

}

func updateMusician(c *gin.Context) {

}

func deleteMusician(c *gin.Context) {

}
