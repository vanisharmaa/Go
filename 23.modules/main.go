package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello Mod in Golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func greeter() {
	fmt.Println("Hey there mod users!")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang</h1>"))
}

// go mod tidy
// go mod veridy
// go list
// go list all
// go list -m all
// go mod why <package>
// go mod graph
// go mod edit -go 1.17.0
// go mod vendor - rather than web, everything is gotten from vendor folder like node modules basically > go run -mod=vendor main.go
