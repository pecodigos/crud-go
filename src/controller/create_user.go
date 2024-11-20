package controller

import (
	"crud-go/src/controller/model/request"
	"crud-go/src/rest_err"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Incorrect field, error=%s", err.Error))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
