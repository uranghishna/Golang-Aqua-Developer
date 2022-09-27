package routes

import (
	"gorm/handler"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, userHandler *handler.UserHandler) {
	e.POST("/createUsers", userHandler.CreateUser)
	e.GET("/users", userHandler.GetAllUser)
	// e.Get("/users/:id", userHandler.GetUserByID)
	e.PUT("/updateUsers/:id", userHandler.UpdateUser)
	e.DELETE("/deleteUsers/:id", userHandler.DeleteUser)
}
