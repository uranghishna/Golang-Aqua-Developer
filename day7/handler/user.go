package handler

import (
	"gorm/entity"
	"gorm/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUsecase *usecase.UserUsecase
}

func NewUserHandler(UserUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase}
}

func (handler UserHandler) CreateUser(c echo.Context) error {
	req := entity.UserRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := handler.UserUsecase.CreateUser(req)

	if err != nil {
		return err
	}

	return c.JSON(201, user)
}

func (handler UserHandler) GetAllUser(c echo.Context) error {

	user, err := handler.UserUsecase.GetAllUser()

	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

// func (handler UserHandler) GetUserByID(c echo.Context) error {

// 	user, err := handler.UserUsecase.GetUserByID()

// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(200, user)

// }

func (handler UserHandler) UpdateUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	userRequest := entity.UpdateUserRequest{}
	if err := c.Bind(&userRequest); err != nil {
		return err
	}
	user, err := handler.UserUsecase.UpdateUser(userRequest, userId)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

func (handler UserHandler) DeleteUser(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))
	err := handler.UserUsecase.DeleteUser(userId)
	if err != nil {
		return err
	}
	return c.JSON(200, userId)
}
