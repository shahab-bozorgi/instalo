package main

import (
	"log"

	"github.com/shahab-bozorgi/instalo/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running",
		})
	})
	log.Printf("Server started on port %s", cfg.Port)

	r.Run(":" + cfg.Port)
}
