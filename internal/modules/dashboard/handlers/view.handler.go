package handlers

import (
	layouts "github.com/hsrvms/fixparts/web/templates/layouts/base"
	dashboardview "github.com/hsrvms/fixparts/web/templates/pages/dashboard"
	"github.com/labstack/echo/v4"
)

func ViewHandler(c echo.Context) error {
	component := layouts.Layout(dashboardview.Dashboard())
	return component.Render(c.Request().Context(), c.Response().Writer)
}
