package routes

import (
	"crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserById())
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail())
	r.POST("/createUser", controller.CreateUser())
	r.PUT("/updateUser", controller.UpdateUser())
	r.DELETE("/deleteUser", controller.DeleteUser())
}
