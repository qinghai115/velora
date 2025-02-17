package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

	fmt.Println("Velora API Gateway")
}
