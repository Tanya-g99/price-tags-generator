package handlers

import (
	"backend/models"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
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
	file, fileHeader, err := r.FormFile("storeLogo")
	if err == nil {
		defer file.Close()

		// Читаем первые байты для определения типа файла
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			log.Println("Ошибка чтения файла логотипа:", err)
			logoBase64 = ""
		} else {
			sampleSize := 512
			if len(fileBytes) < sampleSize {
				sampleSize = len(fileBytes)
			}
			contentType := http.DetectContentType(fileBytes[:sampleSize])
			fileName := fileHeader.Filename
			fileExt := strings.ToLower(filepath.Ext(fileName))

			// Определяем MIME-тип по расширению или содержимому
			var mimeType string
			if strings.Contains(contentType, "svg") || fileExt == ".svg" {
				// SVG - обрабатываем как текст
				mimeType = "image/svg+xml"
				logoBase64 = "data:" + mimeType + ";base64," + base64.StdEncoding.EncodeToString(fileBytes)
			} else {
				// Растровые изображения - декодируем и кодируем в исходном формате
				reader := bytes.NewReader(fileBytes)
				img, format, err := image.Decode(reader)
				if err != nil {
					log.Println("Ошибка декодирования изображения:", err)
					logoBase64 = ""
				} else {
					var buf bytes.Buffer
					switch format {
					case "jpeg", "jpg":
						mimeType = "image/jpeg"
						err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 90})
					case "png":
						mimeType = "image/png"
						err = png.Encode(&buf, img)
					case "gif":
						mimeType = "image/gif"
						err = gif.Encode(&buf, img, &gif.Options{})
					default:
						// Если формат не поддерживается, конвертируем в PNG
						mimeType = "image/png"
						err = png.Encode(&buf, img)
					}
					if err != nil {
						log.Println("Ошибка кодирования изображения:", err)
						logoBase64 = ""
					} else {
						logoBase64 = "data:" + mimeType + ";base64," + base64.StdEncoding.EncodeToString(buf.Bytes())
					}
				}
			}
		}
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
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	baseDir := filepath.Dir(exePath)

	// если программа запущена из временного каталога (go run)
	if strings.Contains(baseDir, os.TempDir()) {
		wd, _ := os.Getwd()
		baseDir = wd
	}

	uploadDir := filepath.Join(baseDir, "uploads")
	os.MkdirAll(uploadDir, 0755)

	svgFile := fmt.Sprintf("pricetag_%s.svg", nextPriceTagID)
	svgPath := filepath.Join(uploadDir, svgFile)

	err = os.WriteFile(svgPath, []byte(svgContent), 0644)
	if err != nil {
		log.Println("❌ Ошибка при сохранении SVG:", err)
	} else {
		log.Println("✅ SVG сохранён в:", svgPath)
	}

	// путь, который отдаётся клиенту
	pt := models.PriceTag{
		ID:            nextPriceTagID,
		ProductID:     productID,
		ProductName:   productName,
		Price:         productPrice,
		PriceTagImage: "uploads/" + svgFile,
	}

	priceTags = append(priceTags, pt)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pt)
	log.Println("✔ Ценник создан:", pt)
}
