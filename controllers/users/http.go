package controllers

import (
	"net/http"
	"strconv"
	"widyaWicaraBackend/businesses/users"
	"widyaWicaraBackend/controllers"
	"widyaWicaraBackend/controllers/users/request"
	"widyaWicaraBackend/controllers/users/response"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	usecase users.Usecase
}

func NewUserController(uc users.Usecase)*UserController{
	return &UserController{
		usecase:uc,
	}
}


func (controller *UserController) Login(c echo.Context)error{
	ctx := c.Request().Context()

	var UserLogin request.UserLogin
	err := c.Bind(&UserLogin) 

	if err != nil{
		return controllers.NewErrorResponse(c,http.StatusInternalServerError,err)

	}
	user,err := controller.usecase.Login(*UserLogin.ToDomain(),ctx)
	if err != nil{
		return controllers.NewErrorResponse(c,http.StatusInternalServerError,err)
	}

	return controllers.NewSuccessResponse(c,response.FromDomain(user))
}

func (controller *UserController) CreateNewUser(c echo.Context)error{
	ctx := c.Request().Context()

	var CreateUser request.CreateNewUser

	err := c.Bind(&CreateUser)

	if err != nil {
		return controllers.NewErrorResponse(c,http.StatusInternalServerError,err)
	}
	user,err := controller.usecase.CreateNewUser(*CreateUser.ToDomain(),ctx)
	if err != nil{
		return controllers.NewErrorResponse(c,http.StatusInternalServerError,err)
	}
	return controllers.NewSuccessResponse(c,response.FromDomain(user))
}

func (controller *UserController) FindById(c echo.Context)error{
	ctx := c.Request().Context()

	id,_:= strconv.Atoi(c.Param("id"))

	user,err := controller.usecase.FindById(id,ctx)
	
	if err != nil{
		return controllers.NewErrorResponse(c,http.StatusInternalServerError,err)
	}
	return controllers.NewSuccessResponse(c,response.FromDomain(user))
}