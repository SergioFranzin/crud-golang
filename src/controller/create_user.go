package controller

import (
	"net/http"

	"github.com/SergioFranzin/crud-golang/src/configuration/logger"
	"github.com/SergioFranzin/crud-golang/src/configuration/validation"
	"github.com/SergioFranzin/crud-golang/src/model/request"
	"github.com/SergioFranzin/crud-golang/src/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createuser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createuser"),
		)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	logger.Info("User created sucessfully",
		zap.String("journey", "createuser"),
	)

	c.JSON(http.StatusOK, response)
}
