package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	mux.HandleFunc("/example/", ExStatic)

	log.Println("http://localhost:4000/")
	log.Println("http://localhost:4000/snippet/create")
	log.Println("http://localhost:4000/snippet/view")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Cache-Control", "public, max-age=31536000")
	// w.Header().Add("cACHE-CONTROL", "public")
	// w.Header().Add("Cache-Control", "max-age=31536000")
	// // w.Header().Del("Cache-Control")
	// w.Header().Get("Cache-Control")

	// w.Write([]byte("Display a specific snippet... "))

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(405)

		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet... "))
}

func ExStatic(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("example"))
}
