package repository_mysql

import (
	"context"
	"go-base/model"
	"go-base/request"
	"log"

	"gorm.io/gorm"
)

type IUserRepo interface {
	GetUsers(ctx context.Context) (users []model.User, err error)
	CreateUser(ctx context.Context, request request.PostUser) (user model.User, err error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) IUserRepo {
	return &UserRepo{DB}
}

func (repo *UserRepo) GetUsers(ctx context.Context) (users []model.User, err error) {
	if err = repo.DB.Debug().WithContext(ctx).Find(&users).Error; err != nil {
		log.Printf("Failed Find With Error : %v", err)
		return
	}

	return
}

func (repo *UserRepo) CreateUser(ctx context.Context, request request.PostUser) (user model.User, err error) {
	user = model.User{
		Name:     request.Name,
		Fullname: request.Fullname,
	}

	if err = repo.DB.Debug().WithContext(ctx).Create(&user).Error; err != nil {
		log.Printf("Failed Create With Error : %v", err)
		return
	}

	return
}
