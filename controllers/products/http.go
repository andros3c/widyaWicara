package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"widyaWicaraBackend/businesses/products"
	"widyaWicaraBackend/controllers"
	"widyaWicaraBackend/controllers/products/request"
	"widyaWicaraBackend/controllers/products/response"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	usecase products.Usecase
}

func NewProductController(uc products.Usecase)*ProductController{
	return &ProductController{
		usecase: uc,
	}
}

func (controller *ProductController) AddProduct(c echo.Context)error{
	ctx := c.Request().Context()

	var AddProd request.Product
	err := c.Bind(&AddProd)

	if err != nil{
		return controllers.NewErrorResponse(c,http.StatusInternalServerError,err)

	}
	prod,err := controller.usecase.AddProduct(*AddProd.ToDomain(),ctx)
	if err != nil{
		return controllers.NewErrorResponse(c,http.StatusInternalServerError,err)
	}

	return controllers.NewSuccessResponse(c,response.FromDomain(prod))

}

func (controller *ProductController) ShowAll(c echo.Context)error{
	ctx := c.Request().Context()

	user_id,_:= strconv.Atoi(c.Param("user_id"))

	prod,err := controller.usecase.ShowAll(user_id,ctx)

	if err != nil{
		return controllers.NewErrorResponse(c,http.StatusInternalServerError,err)
	}
	respon := []response.ProductResponse{}
	for _, value := range prod {
		respon = append(respon,response.FromDomain(value))
	}

	return controllers.NewSuccessResponse(c,respon)

}
func (controller *ProductController) FindById(c echo.Context)error{
	ctx := c.Request().Context()

	id,_:= strconv.Atoi(c.Param("id"))
	prod,err := controller.usecase.FindById(id,ctx)

	if err != nil{
		return controllers.NewErrorResponse(c,http.StatusInternalServerError,err)
	}

	return controllers.NewSuccessResponse(c,response.FromDomain(prod))

}
func (controller *ProductController) Delete(c echo.Context)error{
	ctx := c.Request().Context()

	id,_:= strconv.Atoi(c.Param("id"))
	prod,err := controller.usecase.Delete(id,ctx)

	if err != nil{
		return controllers.NewErrorResponse(c,http.StatusInternalServerError,err)
	}

	return controllers.NewSuccessResponse(c,response.FromDomain(prod))

}

func (controller *ProductController) Update(c echo.Context)error{
	ctx := c.Request().Context()
	
	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}



	id,_:= strconv.Atoi(idstr)
	req := request.Product{}
	if err := c.Bind(&req); 
	err != nil{
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	
	}
	domainReq := req.ToDomain()

	
	
	prod,err := controller.usecase.Update(id,*domainReq,ctx)

	if err != nil{
		return controllers.NewErrorResponse(c,http.StatusInternalServerError,err)
	}

	return controllers.NewSuccessResponse(c,response.FromDomain(prod))

}