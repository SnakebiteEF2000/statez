package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SnakebiteEF2000/statez"
)

type natsService struct {
	name          string
	statezService *statez.Service
}

func main() {

	sz := statez.NewStatez("example2")

	natsSvc := natsService{
		name:          "natsService",
		statezService: statez.NewService("natsService"),
	}

	sz.RegisterService(natsSvc.statezService)

	// go natsSvc.Run or do action

	fmt.Println(natsSvc.statezService.GetState())

	go func(n natsService) {
		time.Sleep(5 * time.Second)
		natsSvc.statezService.StateReady()
		fmt.Println(natsSvc.statezService.GetState())

	}(natsSvc)

	http.HandleFunc("/ready", sz.ReadinessHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
