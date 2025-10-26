package services

import (
	"bytes"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

func GenerateBarcode(data string) ([]byte, error) {
	// Генерация штрихкода Code128
	bc, err := code128.Encode(data)
	if err != nil {
		return nil, err
	}

	// Масштабируем сразу без присваивания
	scaled, err := barcode.Scale(bc, 300, 150)
	if err != nil {
		return nil, err
	}

	// Конвертируем в PNG
	var buf bytes.Buffer
	if err := png.Encode(&buf, scaled); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
