package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"goframe/service"
)



func  GetAll(c echo.Context) error {
	return c.JSON(http.StatusCreated,service.Frist())
}

func GetAllUR(c echo.Context) error{
	return c.JSON(http.StatusOK,service.MQuery())
}

func Find(c echo.Context) error{
	return c.JSON(http.StatusOK,service.Find())
}
