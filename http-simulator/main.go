package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func MakeRequest() {
	log.Println("Go Routine")
	rand.Seed(time.Now().UnixNano())
	for {
		n := rand.Intn(10) // n will be between 0 and 10
		time.Sleep(time.Duration(n) * time.Second)
		_, err := http.Get("http://go-service/v1/ping")
		if err != nil {
			log.Println("Not possible to do request")
		}
	}
}

func main() {
	log.Println("Starting up")
	for i := 0; i < 2; i++ {
		go MakeRequest()
	}
	c := make(chan struct{})
	<-c
}
