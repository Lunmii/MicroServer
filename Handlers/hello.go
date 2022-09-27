package Handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *hello {
	return &hello{l}
}

func (h *hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello There!!")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "You're so wrong", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)
}
