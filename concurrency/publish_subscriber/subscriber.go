package main

// Subscriber ...
type Subscriber interface {
	Notify(interface{}) error
	Close()
}
