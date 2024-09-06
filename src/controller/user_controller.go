package controller

import (
	"LearnGo/src/controller/model/request"
	"LearnGo/src/controller/model/response"
	"LearnGo/src/model"
	"LearnGo/src/setting/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	IUserDomain model.IUserDomain
)

func FindUserById(c *gin.Context) {

}

func FindUserByEmail(c *gin.Context) {

}

func CreateUser(c *gin.Context) {
	//logger.Info("Teste")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		//logger.Error("Error trying to marshal object", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}
	var userResponse response.UserResponse

	userResponse.Email = userRequest.Email
	userResponse.Name = userRequest.Name
	userResponse.Age = userRequest.Age

	c.JSON(http.StatusCreated, userResponse)
}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
