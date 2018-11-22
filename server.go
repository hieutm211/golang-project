package main

import ( 
   "fmt"
   "net/http"
   "github.com/gorilla/mux"
   "time"
   "log"
)

var router = mux.NewRouter()

func LocationHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    fmt.Fprintln(w, "Your country:", vars["country"])
    fmt.Fprintln(w, "Your City:", vars["city"])
    fmt.Fprintln(w, "Your Street:", vars["street"])
    fmt.Fprintln(w, "Your Street Number:", vars["streetnum"])
}

func Location2(w http.ResponseWriter, r *http.Request) {
    country := r.FormValue("country")
    city := r.FormValue("city")
    street := r.FormValue("street")
    number := r.FormValue("number")

    url, err := router.Get("location").URL(
	"country", country,
	"city", city,
	"streetnum", number,
	"street", street,
    )
    if err != nil { panic(err) }
    
    fmt.Fprintf(w, "<!DOCTYPE html><html><a href=%v>go</a></html>", url)
}

func main() {
    router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("github.com/hieutm211/golang-project/"))))
    router.HandleFunc("/location/{country}/{city}/{streetnum:[0-9]+}{street}", LocationHandler).Name("location")
    router.HandleFunc("/enter", Location2)

    server := &http.Server{
	Handler: router,
	Addr: ":8080",
	WriteTimeout: 15*time.Second,
	ReadTimeout: 15*time.Second,
    }


    log.Fatal(server.ListenAndServe())
}
