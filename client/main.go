package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func SendFiles(archive string) {
	file, _ := os.Open(archive)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	r, _ := http.NewRequest("POST", "http://127.0.0.1:8000/upload", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
        client := &http.Client{}
	client.Do(r)
}

func main() {
	SendFiles("./static.tar.gz")
}
