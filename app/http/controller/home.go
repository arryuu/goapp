package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Home struct {
}

func (*Home) GET(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"time": time.Now().Format("2006-01-02 15:04:05"),
	})
}
