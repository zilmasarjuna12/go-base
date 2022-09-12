package main

import (
	"go-base/controller"
	"go-base/migrations"
	repository_mysql "go-base/repository/mysql"
	"go-base/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := repository_mysql.NewRepositories()
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}

	// running migrations
	migrations.Migration()

	userRepo := repository_mysql.NewUserRepo(db)

	userUsecase := service.NewUserUsecase(userRepo)

	userController := controller.NewUserHandler(userUsecase)

	app := gin.New()
	app.Use(gin.Logger())

	v1 := app.Group("v1")
	{
		user := v1.Group("user")
		{
			user.GET("", userController.GetUsers)
			user.POST("", userController.PostUser)
		}
	}

	app.Run()
}
