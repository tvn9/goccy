package main

import (
	"log"
	"net/http"
)

func Producer() <-chan string {
	c := make(chan string)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Recovered", r)
			}
		}()
		for i := 0; i < 10; i++ {
			c <- produceValue()
		}
		close(c)
	}()
	return c
}

func httpHandler(w http.ResponseWriter, req *http.Request) {
	c := Producer()
	for _, x := range c {
		w.Write([]bytes(x))
	}
}
