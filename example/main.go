package main

import (
	"log"
	"net/http"

	"github.com/SnakebiteEF2000/statez"
)

func main() {
	sz := statez.Statez{}

	testSvc := statez.NewServiceHandlerWithOpts("mqtt handler")
	testSvc.StateReady()

	test2Svc := statez.NewServiceHandlerWithOpts("web server")
	test2Svc.StateReady()

	test3Svc := statez.NewServiceHandlerWithOpts("database")
	test3Svc.StateIgnore()

	sz.RegisterService(testSvc, test2Svc, test3Svc)

	http.HandleFunc("/ready", sz.ReadynessHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
