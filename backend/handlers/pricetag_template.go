package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func ListPriceTagTemplates(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("templates/pricetags")
	fmt.Println("ListPriceTagTemplates 1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("ListPriceTagTemplates 2")

	templates := []string{}
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".svg" {
			templates = append(templates, f.Name())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(templates)
}

func AddPriceTagTemplate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Begin AddPriceTagTemplate")
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Ошибка при разборе формы: "+err.Error(), http.StatusBadRequest)
		fmt.Println("1")
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Файл не найден: "+err.Error(), http.StatusBadRequest)
		fmt.Println("2")
		return
	}
	defer file.Close()

	if filepath.Ext(header.Filename) != ".svg" {
		http.Error(w, "Разрешены только .svg файлы", http.StatusBadRequest)
		fmt.Println("3", filepath.Ext(header.Filename))
		return
	}

	dstPath := filepath.Join("templates/pricetags", header.Filename)

	dst, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, "Не удалось создать файл: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Копируем содержимое
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Ошибка при сохранении: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Успешный ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":   "ok",
		"filename": header.Filename,
	})
}
