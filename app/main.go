package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Message from sample-app"))
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
