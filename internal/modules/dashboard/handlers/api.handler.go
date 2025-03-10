package handlers

import (
	"fmt"

	"github.com/hsrvms/fixparts/internal/modules/dashboard/services"
	"github.com/labstack/echo/v4"
)

type DashboardAPIHandler struct {
	service services.DashboardService
}

func NewDashboardAPIHandler(service services.DashboardService) *DashboardAPIHandler {
	return &DashboardAPIHandler{service: service}
}

func (h *DashboardAPIHandler) GetLowStockCount(c echo.Context) error {
	ctx := c.Request().Context()
	count, err := h.service.GetLowStockCount(ctx)
	if err != nil {
		return c.HTML(500, "<div class='text-red-600'>Düşük stok sayısı alınamadı</div>")
	}
	return c.HTML(200, fmt.Sprintf("<div>%d</div>", count))
}

func (h *DashboardAPIHandler) GetTodaySales(c echo.Context) error {
	ctx := c.Request().Context()
	sales, err := h.service.GetTodaySales(ctx)
	if err != nil {
		return c.HTML(500, "<div class='text-red-600'>Bugünkü satışlar alınamadı</div>")
	}
	return c.HTML(200, fmt.Sprintf("<div>%.2f</div>", sales))
}

func (h *DashboardAPIHandler) GetTotalInventoryCount(c echo.Context) error {
	ctx := c.Request().Context()
	count, err := h.service.GetTotalInventoryCount(ctx)
	if err != nil {
		return c.HTML(500, "<div class='text-red-600'>Envanter sayısı alınamadı</div>")
	}
	return c.HTML(200, fmt.Sprintf("<div>%d</div>", count))
}

func (h *DashboardAPIHandler) GetVehicleCount(c echo.Context) error {
	ctx := c.Request().Context()
	count, err := h.service.GetVehicleCount(ctx)
	if err != nil {
		return c.HTML(500, "<div class='text-red-600'>Araç sayısı alınamadı</div>")
	}
	return c.HTML(200, fmt.Sprintf("<div>%d</div>", count))
}

func (h *DashboardAPIHandler) GetLowStockItems(c echo.Context) error {
	ctx := c.Request().Context()
	items, err := h.service.GetLowStockItems(ctx)

	if err != nil {
		return c.HTML(500, "<div class='text-red-600'>Düşük stoklu ürünler alınamadı</div>")
	}

	if len(items) == 0 {
		return c.HTML(200, "<div class='text-gray-500'>Düşük stoklu ürün bulunamadı</div>")
	}

	html := "<table class='min-w-full'><thead><tr>" +
		"<th>Parça Numarası</th><th>İsim</th><th>Mevcut Stok</th><th>Minimum Stok</th>" +
		"</tr></thead><tbody>"

	for _, item := range items {
		html += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%d</td><td>%d</td></tr>",
			item.PartNumber, item.Name, item.Current, item.Minimum)
	}
	html += "</tbody></table>"

	return c.HTML(200, html)
}

func (h *DashboardAPIHandler) GetRecentSales(c echo.Context) error {
	ctx := c.Request().Context()
	sales, err := h.service.GetRecentSales(ctx)
	if err != nil {
		return c.HTML(500, "<div class='text-red-600'>Son satışlar alınamadı</div>")
	}

	if len(sales) == 0 {
		return c.HTML(200, "<div class='text-gray-500'>Son satış bulunamadı</div>")
	}

	html := "<table class='min-w-full'><thead><tr>" +
		"<th>Tarih</th><th>Parça</th><th>Müşteri</th><th>Toplam</th>" +
		"</tr></thead><tbody>"

	for _, sale := range sales {
		html += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td><td>%.2f</td></tr>",
			sale.Date, sale.Part, sale.Customer, sale.Total)
	}
	html += "</tbody></table>"

	return c.HTML(200, html)
}

func (h *DashboardAPIHandler) GetTopSellers(c echo.Context) error {
	ctx := c.Request().Context()
	sellers, err := h.service.GetTopSellers(ctx)
	if err != nil {
		return c.HTML(500, "<div class='text-red-600'>En çok satanlar alınamadı</div>")
	}

	if len(sellers) == 0 {
		return c.HTML(200, "<div class='text-gray-500'>En çok satan ürün bulunamadı</div>")
	}

	html := "<table class='min-w-full'><thead><tr>" +
		"<th>Parça Numarası</th><th>İsim</th><th>Satılan</th><th>Gelir</th>" +
		"</tr></thead><tbody>"

	for _, seller := range sellers {
		html += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%d</td><td>%.2f</td></tr>",
			seller.PartNumber, seller.Name, seller.Sold, seller.Revenue)
	}
	html += "</tbody></table>"

	return c.HTML(200, html)
}

func (h *DashboardAPIHandler) GetRecentPurchases(c echo.Context) error {
	ctx := c.Request().Context()
	purchases, err := h.service.GetRecentPurchases(ctx)
	if err != nil {
		return c.HTML(500, "<div class='text-red-600'>Son alımlar alınamadı</div>")
	}

	if len(purchases) == 0 {
		return c.HTML(200, "<div class='text-gray-500'>Son alım bulunamadı</div>")
	}

	html := "<table class='min-w-full'><thead><tr>" +
		"<th>Tarih</th><th>Parça Numarası</th><th>Tedarikçi</th><th>Maliyet</th>" +
		"</tr></thead><tbody>"

	for _, purchase := range purchases {
		html += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td><td>%.2f</td></tr>",
			purchase.Date, purchase.PartNumber, purchase.Supplier, purchase.Cost)
	}
	html += "</tbody></table>"

	return c.HTML(200, html)
}
