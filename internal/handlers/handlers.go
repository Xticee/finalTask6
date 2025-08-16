package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func GetForm(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs("../index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func UploadForm(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := service.Decode(string(fileContent))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = os.WriteFile("../result.txt", []byte(resp), 0755)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
