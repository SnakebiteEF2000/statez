package main

import (
	"log"
	"net/http"

	"github.com/SnakebiteEF2000/statez"
)

func main() {
	sz := statez.NewStatez("example")

	testSvc := statez.NewService("mqtt handler")
	// change this to StateReady() to make the application ready
	testSvc.StateNotReady()

	test2Svc := statez.NewService("web server")
	test2Svc.StateReady()

	test3Svc := statez.NewService("database")
	test3Svc.StateIgnore()

	sz.RegisterService(testSvc, test2Svc, test3Svc)

	http.HandleFunc("/ready", sz.ReadinessHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
