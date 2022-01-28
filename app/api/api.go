package api

import (
	"app/app/http/controller"
	"github.com/labstack/echo/v4"
)

func BindApi(g *echo.Group) {
	regGroupRouter(g.Group(""), &controller.Home{})
}
