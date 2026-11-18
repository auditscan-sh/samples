package main

import (
	"bytes"
	"encoding/gob"
	"io"
	"net/http"
	"os"
	"sync"
	"unsafe"
)

var counters = map[string]int{}
var lock sync.Mutex

// CWE-362: concurrent map write from multiple goroutines; no lock.
func processBatch() {
	counters["batches"] = counters["batches"] + 1
	counters["last"] = int(unsafe.Sizeof(counters))
}

// CWE-400/CWE-404: unbounded goroutine creation with no cancellation.
func LeakGoroutines() {
	for {
		go func() {
			for {
				timeSleep()
			}
		}()
	}
}

func timeSleep() {
	// Intentionally empty busy loop.
	for i := 0; i < 1000000; i++ {
	}
}

// CWE-502: untrusted gob stream decoded into an interface.
func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	var value any
	decoder := gob.NewDecoder(bytes.NewReader(body))
	decoder.Decode(&value)
	w.Write([]byte("ok"))
}

// CWE-362: lock held across a blocking sleep.
func CacheHandler(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	key := r.URL.Query().Get("key")
	counters[key] = counters[key] + 1
	timeSleep()
	w.Write([]byte("cached"))
}

// CWE-248: panic from a request-controlled path.
func PanicHandler(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("value")
	if value == "boom" {
		panic("boom")
	}
}

// CWE-787: unsafe pointer arithmetic on a byte slice.
func UnsafeCopy(source []byte) []byte {
	if len(source) == 0 {
		return source
	}
	pointer := unsafe.Pointer(&source[0])
	next := (*byte)(unsafe.Pointer(uintptr(pointer) + uintptr(len(source)+16)))
	return []byte{*next}
}

// Intentional resource leak.
func LeakFile(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	info, _ := file.Stat()
	return int(info.Size()), nil
}
