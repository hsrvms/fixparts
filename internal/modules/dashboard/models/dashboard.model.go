package models

type LowStockItem struct {
	PartNumber string `json:"part_number"`
	Name       string `json:"name"`
	Current    int    `json:"current"`
	Minimum    int    `json:"minimum"`
}

type RecentSale struct {
	Date     string  `json:"date"`
	Part     string  `json:"part"`
	Customer string  `json:"customer"`
	Total    float64 `json:"total"`
}

type TopSeller struct {
	PartNumber string  `json:"part_number"`
	Name       string  `json:"name"`
	Sold       int     `json:"sold"`
	Revenue    float64 `json:"revenue"`
}

type RecentPurchase struct {
	Date       string  `json:"date"`
	PartNumber string  `json:"part_number"`
	Supplier   string  `json:"supplier"`
	Cost       float64 `json:"cost"`
}
