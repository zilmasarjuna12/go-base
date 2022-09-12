package controller

import (
	"go-base/request"
	"go-base/response"
	"go-base/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase service.IUserUsecase
}

func NewUserHandler(
	userUsecase service.IUserUsecase,
) *UserHandler {
	return &UserHandler{userUsecase}
}

func (handler *UserHandler) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()

	res, statusCode, err := handler.userUsecase.GetUsers(ctx)

	if err != nil {
		log.Printf("Failed handler.userUsecase.GetUsers Got Error %v", err)
		c.JSON(statusCode, res)
		return
	}

	c.JSON(statusCode, res)
}

func (handler *UserHandler) PostUser(c *gin.Context) {
	ctx := c.Request.Context()

	var req request.PostUser

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed ShouldBindJSON %v", err)
		res := response.BaseResponse{}
		response.NewBaseResponseStatusCode(http.StatusBadRequest, &res, err)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	res, statusCode, err := handler.userUsecase.PostUser(ctx, req)

	if err != nil {
		log.Printf("Failed handler.userUsecase.PostUser Got Error %v", err)
		c.JSON(statusCode, res)
		return
	}

	c.JSON(statusCode, res)
}
