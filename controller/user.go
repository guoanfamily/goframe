package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"goframe/service"
)



func  GetAll(c echo.Context) error {
	return c.JSON(http.StatusCreated,service.Frist())
}
