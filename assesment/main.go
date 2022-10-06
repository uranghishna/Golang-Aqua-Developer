package main

import (
	"ecommerce/config"
	"ecommerce/docs"
	"ecommerce/handler"
	"ecommerce/repository"
	"ecommerce/routes"
	"ecommerce/usecase"

	"github.com/labstack/echo/v4"
	// _ "github.com/swaggo/echo-swagger/example/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io

func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "E-Commerce"
	docs.SwaggerInfo.Description = "Assessment Associate Software Engineer eFishery Academy Batch 2.0"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.Getenv("SWAGGER_HOST", "localhost:8000")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	config.Database()
	// config.Migrate()

	e := echo.New()
	e.Static("/storage", "storage/")

	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	productRepository := repository.NewProductRepository(config.DB)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productHandler := handler.NewProductHandler(productUsecase)

	cartRepository := repository.NewCartRepository(config.DB)
	cartUsecase := usecase.NewCartUsecase(cartRepository)
	cartHandler := handler.NewCartHandler(cartUsecase)

	routes.Routes(e, userHandler, productHandler, cartHandler)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8000"))
}
