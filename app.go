package main

import (
	"log"
	"net/http"

	"field/ctl"
)

func main() {
	http.HandleFunc("/", ctl.Health)
	http.HandleFunc("/field", ctl.Fields)
	http.HandleFunc("/field/update", ctl.FieldUpdate)

	log.Println("server start in 9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}
