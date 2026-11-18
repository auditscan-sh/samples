package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const uploadDir = "/srv/uploads"

// CWE-434 + CWE-22: client filename, no type check, no size cap.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(64 << 20)
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	target := filepath.Join(uploadDir, header.Filename)
	out, err := os.Create(target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()
	io.Copy(out, file)
	w.Write([]byte("saved " + target))
}

// CWE-22: traversal in ServeFile.
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	http.ServeFile(w, r, filepath.Join(uploadDir, name))
}

// CWE-22/CWE-73: request value selects the deleted path.
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	os.Remove(filepath.Join(uploadDir, name))
	w.Write([]byte("deleted"))
}

// CWE-22: archive extraction does not validate member paths.
func ExtractHandler(w http.ResponseWriter, r *http.Request) {
	archive := r.FormValue("archive")
	destination := r.FormValue("destination")
	ExtractZip(archive, destination)
	w.Write([]byte("extracted"))
}

// CWE-918 chain: handler -> FetchPreview -> http.Get(target).
func PreviewHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")
	body, err := FetchPreview(target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.Write([]byte(body))
}

// CWE-78: request file reaches a shell command in the same handler.
func ThumbnailHandler(w http.ResponseWriter, r *http.Request) {
	file := r.FormValue("file")
	command := exec.Command("sh", "-c", "convert "+file+" -resize 200x200 /srv/uploads/thumb.png")
	output, err := command.CombinedOutput()
	if err != nil {
		http.Error(w, string(output), http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

// CWE-918: request URL is fetched by the server.
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("callback")
	body, err := FetchPreview(target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.Write([]byte(body))
}

// CWE-400: unbounded read.
func ReadAllHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	w.Write(body)
}
