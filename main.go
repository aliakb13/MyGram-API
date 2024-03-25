package main

import (
	"final-project/config"
	"final-project/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.StartDB()

	ginEngine := gin.Default()

	router.StartRouter(ginEngine, db)

	err := ginEngine.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
