package routes

import (
	"ecommerce/handler"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, userHandler *handler.UserHandler, productHandler *handler.ProductHandler) {
	e.POST("/users/create", userHandler.CreateUser)
	e.GET("/users", userHandler.GetAllUser)
	// e.Get("/users/:id", userHandler.GetUserByID)
	e.PUT("/users/update/:id", userHandler.UpdateUser)
	e.DELETE("/users/delete/:id", userHandler.DeleteUser)

	e.POST("/products/create", productHandler.CreateProduct)
	e.GET("/products", productHandler.GetAllProduct)
	// e.Get("/products/:id", productHandler.GetProductByID)
	e.PUT("/products/update/:id", productHandler.UpdateProduct)
	e.DELETE("/products/delete/:id", productHandler.DeleteProduct)

	e.Static("/static", "storage")
}
