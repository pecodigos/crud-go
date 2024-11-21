package main

import (
	"crud-go/src/configuration/logger"
	"crud-go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	logger.Info("About to start app.")
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
