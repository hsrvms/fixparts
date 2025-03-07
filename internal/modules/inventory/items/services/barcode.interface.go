package services

type BarcodeService interface {
	GenerateBarcode(categoryID int, supplierID int) (string, error)
	GenerateBarcodeImage(barcodeText string) ([]byte, error)
}
