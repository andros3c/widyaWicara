package routes

import (
	ProductController "widyaWicaraBackend/controllers/products"
	UserController "widyaWicaraBackend/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


type ControllerList struct{
	UserController 		UserController.UserController
	ProductController	ProductController.ProductController
	JWTConfig				middleware.JWTConfig
}

func (cl *ControllerList) RouteRegister(e *echo.Echo){
	users := e.Group("user")
	users.POST("/login",cl.UserController.Login)
	users.POST("/create",cl.UserController.CreateNewUser)
	users.GET("/:id",cl.UserController.FindById,middleware.JWTWithConfig(cl.JWTConfig))

	products := e.Group("products")
	products.POST("/add",cl.ProductController.AddProduct,middleware.JWTWithConfig(cl.JWTConfig))
	products.GET("/:user_id",cl.ProductController.ShowAll,middleware.JWTWithConfig(cl.JWTConfig))
	products.GET("/find/:id",cl.ProductController.FindById,middleware.JWTWithConfig(cl.JWTConfig))
	products.PUT("/update/:id",cl.ProductController.Update,middleware.JWTWithConfig(cl.JWTConfig))
	products.DELETE("/delete/:id",cl.ProductController.Delete,middleware.JWTWithConfig(cl.JWTConfig))

}