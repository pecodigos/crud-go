package main

import (
	"crud-go/src/configuration/logger"
	"crud-go/src/controller"
	"crud-go/src/controller/routes"
	service2 "crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	logger.Info("About to start app.")

	service := service2.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
