package controllers

import (
	"fmt"
	"net/http"
)


func RemoveItemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}