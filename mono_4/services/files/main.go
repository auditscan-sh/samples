package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", UploadHandler)
	mux.HandleFunc("/download", DownloadHandler)
	mux.HandleFunc("/delete", DeleteHandler)
	mux.HandleFunc("/extract", ExtractHandler)
	mux.HandleFunc("/preview", PreviewHandler)
	mux.HandleFunc("/thumbnail", ThumbnailHandler)
	mux.HandleFunc("/callback", CallbackHandler)
	http.ListenAndServe("0.0.0.0:8080", mux)
}
