package main

import (
	"gorm/usecase"
	"gorm/repository"
	"gorm/config"
	"gorm/routes"
	"gorm/handler"
	// "gorm/entity"
	"github.com/labstack/echo/v4"
)

func main()  {
	config.Database()
	config.Migrate()

	e := echo.New()
	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	routes.Routes(e, userHandler)
	e.Logger.Fatal(e.Start(":8080"))
}