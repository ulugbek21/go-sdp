package main

import (
	"fmt"
	"log"
	"sync"
)

// RequestHandler ...
type RequestHandler func(i interface{})

// Request ...
type Request struct {
	Data    interface{}
	handler RequestHandler
}

// NewStringRequest ...
func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {
	myRequest := Request{
		Data: s, handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(s)
		},
	}

	return myRequest
}
