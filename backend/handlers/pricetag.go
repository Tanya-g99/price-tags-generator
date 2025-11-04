package handlers

import (
	"backend/models"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

var priceTags []models.PriceTag

func ListPriceTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(priceTags)
}

func AddPriceTag(w http.ResponseWriter, r *http.Request) {
	log.Println("📥 Создание ценника")
	nextPriceTagID := uuid.New().String()

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Ошибка парсинга формы: "+err.Error(), http.StatusBadRequest)
		return
	}

	productID := r.FormValue("productId")
	templateFile := r.FormValue("template")

	var productName string
	var productPrice float64
	for _, p := range products {
		if p.ID == productID {
			productName = p.Name
			productPrice = p.Price
			break
		}
	}

	// Читаем SVG-шаблон
	svgBytes, err := os.ReadFile(filepath.Join("templates/pricetags", templateFile))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Шаблон не найден", http.StatusBadRequest)
		return
	}
	svgContent := string(svgBytes)

	// Подставляем данные
	svgContent = strings.ReplaceAll(svgContent, "Название продукта", productName)
	svgContent = strings.ReplaceAll(svgContent, "Цена", fmt.Sprintf("%.2f", productPrice))

	// ====== ЛОГОТИП ======
	var logoBase64 string
	file, _, err := r.FormFile("storeLogo")
	if err == nil {
		defer file.Close()
		var buf bytes.Buffer
		img, _, _ := image.Decode(file)
		png.Encode(&buf, img)
		logoBase64 = "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())
	} else {
		logoBase64 = "" // можно поставить placeholder
	}
	svgContent = strings.ReplaceAll(svgContent, "/templates/assets/logo.svg", logoBase64)

	// ====== BARCODE ======
	fmt.Println(nextPriceTagID)
	barcodeData := fmt.Sprintf("%s-%.2f", productName, productPrice)
	pngBytes, _ := qrcode.Encode(barcodeData, qrcode.Medium, 200)
	qrBase64 := "data:image/png;base64," + base64.StdEncoding.EncodeToString(pngBytes)
	svgContent = strings.ReplaceAll(svgContent, "/templates/assets/qr-code.svg", qrBase64)

	// Сохраняем SVG во временный файл
	svgPath := filepath.Join("uploads", fmt.Sprintf("pricetag_%s.svg", nextPriceTagID))
	os.WriteFile(svgPath, []byte(svgContent), 0644)
	// fmt.Println(svgContent)

	pt := models.PriceTag{
		ID:            nextPriceTagID,
		ProductID:     productID,
		ProductName:   productName,
		Price:         productPrice,
		PriceTagImage: svgPath,
	}

	priceTags = append(priceTags, pt)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pt)
	log.Println("✔ Ценник создан:", pt)
}
