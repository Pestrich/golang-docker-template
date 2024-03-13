package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	helloHandler := func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprintln(w, "Привет мир!")
	}

	r.HandleFunc("/hello", helloHandler)

	log.Println("Сервер запущен на 80 порту")

	log.Fatal(http.ListenAndServe(":80", r))
}
