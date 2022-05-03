package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/slideclick/lispGol/stack"
)

func echo(w http.ResponseWriter, r *http.Request) {
	data, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("\n%s", string(data))
	fmt.Fprintln(w, string(data))
}

func main() {
	stack.Test("haha")
	http.HandleFunc("/", echo)
	panic(http.ListenAndServe(":8080", nil))
}
