package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"goframe/service"
)

type User struct {
	Name  string //`json:"name" xml:"name" form:"name" query:"name"`
	Email string //`json:"email" xml:"email" form:"email" query:"email"`
}

func  GetAll(c echo.Context) error {
	return c.JSON(http.StatusCreated,service.Frist())
}
