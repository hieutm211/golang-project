
package main

import (
    "fmt"
    "net/http"
    "log"
    "time"
    "github.com/gorilla/mux"
)

var router = mux.NewRouter()

func middleWare(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Before request phase...\n")
	handler.ServeHTTP(w, r)
	fmt.Fprintf(w, "After response phase...")
    })
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "executing main handler...\n")
}

func main() {
    router.Handle("/test", middleWare(http.HandlerFunc(mainHandler)))
    server := &http.Server {
	Handler: router,
	Addr: ":8080",
	ReadTimeout: 15*time.Second,
	WriteTimeout: 15*time.Second,
    }
    log.Fatal(server.ListenAndServe()) 
}
