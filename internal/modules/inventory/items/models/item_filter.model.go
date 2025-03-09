package models

type ItemFilter struct {
	CategoryID *int    `query:"category_id"`
	SupplierID *int    `query:"supplier_id"`
	PartNumber *string `query:"part_number"`
	SearchTerm *string `query:"search"`
	LowStock   *bool   `query:"low_stock"`
	MakeID     *int    `query:"make_id"`
	ModelID    *int    `query:"model_id"`
	SubmodelID *int    `query:"submodel_id"`
	IsActive   *bool   `query:"is_active"`
}
