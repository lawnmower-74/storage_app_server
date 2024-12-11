package main

import (
	"github.com/lawnmower-74/storage_app_server/controller"
	"github.com/lawnmower-74/storage_app_server/db"
	"github.com/lawnmower-74/storage_app_server/repository"
	"github.com/lawnmower-74/storage_app_server/router"
	"github.com/lawnmower-74/storage_app_server/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
