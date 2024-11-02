package main

import (
	"fmt"
	"net/http"
	"os"
)

// MethodGet mengecek apakah request method adalah GET.
func MethodGet(r *http.Request) error {
	if r.Method != http.MethodGet {
		return fmt.Errorf("Method not allowed")
	}
	return nil
}

// CheckDataRequest mengecek apakah parameter data ada di URL.
func CheckDataRequest(r *http.Request) error {
	data := r.URL.Query().Get("data")
	if len(data) == 0 {
		return fmt.Errorf("Data not found")
	}
	return nil
}

// CheckOpenFile mengecek apakah file dengan nama filename bisa dibuka.
func CheckOpenFile(r *http.Request) error {
	filename := r.URL.Query().Get("filename")
	if len(filename) == 0 {
		return fmt.Errorf("File not found")
	}
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("File not found")
	}
	file.Close()
	return nil
}

// MethodHandler menghandle request berdasarkan method GET atau selainnya.
func MethodHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := MethodGet(r)
		if err != nil {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method not allowed")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Method handler passed")
	}
}

// DataHandler menghandle request jika parameter `data` ada atau tidak di URL.
func DataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := CheckDataRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Data not found")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Data handler passed")
	}
}

// OpenFileHandler menghandle request jika parameter `filename` bisa dibuka atau tidak.
func OpenFileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := CheckOpenFile(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "File not found")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Error handler passed")
	}
}

func main() {
	http.HandleFunc("/method", MethodHandler())
	http.HandleFunc("/data", DataHandler())
	http.HandleFunc("/openfile", OpenFileHandler())

	http.ListenAndServe("localhost:8080", nil)
}
