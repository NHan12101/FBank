package main

import (
	"fbank-server/internal/config"
	"fbank-server/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	database.ConnectMySQL(cfg)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Status(200)
	})

	router.Run(":" + cfg.ServerPort)
}
