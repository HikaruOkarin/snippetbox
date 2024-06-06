package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}
	w.Write([]byte("Display the home page"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	// headers := http.Header{}
	w.Write([]byte("Display a specific snippet"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		// In contrast, the Add() method appends a new "Cache-Control" header and can
		// be called multiple times.
		w.Header().Add("Cache-Control", "public")
		w.Header().Add("Cache-Control", "max-age=31536000")
		// Delete all values for the "Cache-Control" header.
		w.Header().Del("Cache-Control")
		// Retrieve the first value for the "Cache-Control" header.
		data := w.Header().Get("Cache-Control")
		fmt.Println(data)
		// w.Header().Set("Allow", "POST")
		// w.Header().Add("Date", "baribir") // что бы изменить значения уже готового хедера
		// content := w.Header().Get("Date") // получить значения хедера

		// fmt.Println(content)
		// w.Header().Del("Allow")
		http.Error(w, "Method Not Allowed", 404) // w.WriteHeader(405) => использовать тока один раз && w.Write([]byte("some text"))
		return
	}

	fmt.Fprintln(w, "Create a new snippet")
}

// PAGE 41
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	log.Println("starting server in http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
