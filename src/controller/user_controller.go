package controller

import (
	"LearnGo/src/controller/model/request"
	"LearnGo/src/setting/validation"
	"github.com/gin-gonic/gin"
)

func FindUserById(c *gin.Context) {

}

func FindUserByEmail(c *gin.Context) {

}

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}
}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
