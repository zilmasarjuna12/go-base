package service

import (
	"context"
	repository_mysql "go-base/repository/mysql"
	"go-base/request"
	"go-base/response"
	"log"
	"net/http"
)

type UserUsecase struct {
	userRepo repository_mysql.IUserRepo
}

type IUserUsecase interface {
	GetUsers(ctx context.Context) (result response.GetUser, statusCode int, err error)
	PostUser(ctx context.Context, req request.PostUser) (result response.PostUser, statusCode int, err error)
}

func NewUserUsecase(
	userRepo repository_mysql.IUserRepo,
) IUserUsecase {
	return &UserUsecase{userRepo}
}

func (usecase *UserUsecase) GetUsers(ctx context.Context) (result response.GetUser, statusCode int, err error) {
	statusCode = http.StatusInternalServerError

	user, err := usecase.userRepo.GetUsers(ctx)

	if err != nil {
		log.Printf("Got usecase.userRepo.GetUsers Error %v", err)

		statusCode = http.StatusInternalServerError
		response.NewBaseResponseStatusCode(statusCode, &result.BaseResponse, err)

		return
	}

	statusCode = http.StatusOK
	response.NewBaseResponseStatusCode(statusCode, &result.BaseResponse, nil)
	result.User = user

	return
}

func (usecase *UserUsecase) PostUser(ctx context.Context, req request.PostUser) (result response.PostUser, statusCode int, err error) {
	statusCode = http.StatusInternalServerError

	user, err := usecase.userRepo.CreateUser(ctx, req)

	if err != nil {
		log.Printf("Got usecase.userRepo.CreateUser Error %v", err)

		statusCode = http.StatusInternalServerError
		response.NewBaseResponseStatusCode(statusCode, &result.BaseResponse, err)

		return
	}

	statusCode = http.StatusOK
	response.NewBaseResponseStatusCode(statusCode, &result.BaseResponse, nil)
	result.User = user

	return
}
