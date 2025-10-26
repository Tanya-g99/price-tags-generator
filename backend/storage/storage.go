package storage

import (
	"os"
)

// Сохраняем PNG в файл
func SavePNG(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}

// Здесь можно добавить хранение PriceTag/Product в память или базу
