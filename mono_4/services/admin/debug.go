package main

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"unsafe"
)

// CWE-200: environment dump, including secrets.
func EnvHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{
		"env":    os.Environ(),
		"args":   os.Args,
		"wd":     mustGetwd(),
		"uid":    os.Getuid(),
		"gid":    os.Getgid(),
		"secret": os.Getenv("AWS_SECRET"),
	})
}

// CWE-215: profiling endpoints exposed without auth.
func PprofHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	pprof.Lookup("heap").WriteTo(w, 0)
}

// CWE-787-adjacent: unsafe pointer / memory reinterpretation.
func UnsafeHandler(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("value")
	bytes := []byte(value)
	header := (*reflectStringHeader)(unsafe.Pointer(&bytes))
	w.Write([]byte(header.Data))
}

type reflectStringHeader struct {
	Data uintptr
	Len  int
}

func mustGetwd() string {
	wd, _ := os.Getwd()
	return wd
}

func GoroutineCount() int {
	return runtime.NumGoroutine()
}

// CWE-22: log file read from a request value.
func LogHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	data, err := ReadLog(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Write(data)
}
