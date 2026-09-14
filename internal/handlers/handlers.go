package handlers

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleIndex(w http.ResponseWriter, _ *http.Request) {
	html, err := os.OpenFile("index.html", os.O_RDONLY, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer html.Close()

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, html)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // 10 MB
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileName := time.Now().UTC().Format("20060102-150405") + ".txt"
	resultFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resultFile.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := service.Сonvert(string(data))

	resultFile.WriteString(result)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
