package server

import (
	"net/http"

	"github.com/hsrvms/fixparts/internal/modules/dashboard"
	"github.com/hsrvms/fixparts/internal/modules/inventory"
	"github.com/hsrvms/fixparts/internal/modules/purchases"
	"github.com/hsrvms/fixparts/internal/modules/sales"
	"github.com/hsrvms/fixparts/internal/modules/suppliers"
	"github.com/hsrvms/fixparts/internal/modules/vehicles"
	"github.com/labstack/echo/v4"
)

func (s *Server) initRoutes() {
	api := s.Echo.Group("/api")

	api.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	dashboard.RegisterRoutes(s.Echo, api, s.DB)
	inventory.RegisterRoutes(s.Echo, api, s.DB)
	vehicles.RegisterRoutes(s.Echo, api, s.DB)
	suppliers.RegisterRoutes(api, s.DB)
	purchases.RegisterRoutes(api, s.DB)
	sales.RegisterRoutes(api, s.DB)

}
