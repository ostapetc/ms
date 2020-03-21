package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var ID int64 = 0

func handler(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt64(&ID, 1)

	_, err := fmt.Fprint(w, ID)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
