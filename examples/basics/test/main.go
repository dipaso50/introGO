package main

import (
	"fmt"
	"introGO/examples/basics/test/cache"
	"introGO/examples/basics/test/handlers"
	"introGO/examples/basics/test/inmemory"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	tokCacheService cache.TokenCacheService
)

func main() {
	tokCacheService = cache.NewTokenCacheService(inmemory.NewInmemoryCache())

	r := mux.NewRouter()
	r.Handle("/token", handlers.CreateToken(tokCacheService)).Methods("GET")
	r.Handle("/consumeResource", handlers.TokenMiddleware(tokCacheService, handlers.SayHello())).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Starting up on 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
