package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type XxxHandler struct {
}

func HandlerXxx(c echo.Context) error {
	return c.String(http.StatusOK, "Chào mừng bạn đã đến HandlerXxx được viết bởi Nim.")
}
