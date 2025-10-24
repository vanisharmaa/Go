package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vanisharmaa/go/25.mongo-api/router"
)

func main() {
	fmt.Println("mongo api in golang")
	r := router.Router()
	err := http.ListenAndServe(":4000", r)
	if err != nil {
		log.Fatal(err)
	}

}
