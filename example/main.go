package main

import (
	"log"
	"net/http"

	"github.com/SnakebiteEF2000/statez"
)

func main() {
	sz := statez.Statez{}

	testSvc := statez.NewServiceHandlerWithOpts("mqtt handler")
	testSvc.StatusIgnore()

	sz.RegisterService(testSvc)

	http.HandleFunc("/ready", sz.ReadynessHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
