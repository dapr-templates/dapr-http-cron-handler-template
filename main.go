package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.OPTIONS("/run", optionsHandler)
	r.POST("/run", runHandler)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func runHandler(c *gin.Context) {
	// TODO: do something interesting here
	log.Printf("schedule received: %v", time.Now())
	c.JSON(http.StatusOK, nil)
}

func optionsHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
	c.Header("Allow", "POST,OPTIONS")
	c.Header("Content-Type", "application/json")
	c.AbortWithStatus(http.StatusOK)
}
