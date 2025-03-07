package services

import (
	"bytes"
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"image/png"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

type barcodeService struct{}

func NewBarcodeService() BarcodeService {
	return &barcodeService{}
}

func (s *barcodeService) GenerateBarcode(categoryID int, supplierID int) (string, error) {
	timestamp := time.Now().Format("060102150405")

	randomBytes := make([]byte, 6)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}

	random := base32.StdEncoding.EncodeToString(randomBytes)[:8]

	barcode := fmt.Sprintf("C%03d-S%03d-%s-%s",
		categoryID,
		supplierID,
		timestamp,
		random,
	)

	return barcode, nil
}

func (s *barcodeService) GenerateBarcodeImage(barcodeText string) ([]byte, error) {
	bc, err := code128.Encode(barcodeText)
	if err != nil {
		return nil, err
	}

	scaled, err := barcode.Scale(bc, 300, 100)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, scaled); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
