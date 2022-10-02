package main

import (
	"ecommerce/config"
	"ecommerce/handler"
	"ecommerce/repository"
	"ecommerce/routes"
	"ecommerce/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Database()
	// config.Migrate()

	e := echo.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	productRepository := repository.NewProductRepository(config.DB)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productHandler := handler.NewProductHandler(productUsecase)

	routes.Routes(e, userHandler, productHandler)
	e.Logger.Fatal(e.Start(":8000"))
}
