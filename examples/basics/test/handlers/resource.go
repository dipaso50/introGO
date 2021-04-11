package handlers

import (
	"fmt"
	"net/http"
)

func SayHello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "You are in!!!!")
	})
}
