package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	name      string
	age       int
	available bool
}

func main() {
	engine := gin.Default()
	engine.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "kazz",
			"age":  22,
		})
	})
	engine.Run(":3000")
}
