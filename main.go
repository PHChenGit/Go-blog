package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Go-blog/internal/logging"
)

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`my first go blog`))
}

type IndexData struct {
	Title   string
	Content string
}

func index(w http.ResponseWriter, r *http.Request) {
	data := &IndexData{
		Title: "HomePage",
		Content: "My first Gin Project",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	var logger = loggin

	http.HandleFunc("/", test)
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9888", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
