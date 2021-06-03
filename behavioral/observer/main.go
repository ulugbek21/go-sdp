package main

import (
	"container/list"
	"fmt"
)

// Observable ...
type Observable struct {
	subs *list.List
}

// Subscribe ...
func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

// Unsubscribe ...
func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

// Fire ...
func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

// Observer ...
type Observer interface {
	Notify(data interface{})
}

// Person ...
// whenever a person catches a cold,
// a doctor must be called
type Person struct {
	Observable
	Name string
}

// NewPerson ...
func NewPerson(name string) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		Name:       name,
	}
}

// CatchACold ...
func (p *Person) CatchACold() {
	p.Fire(p.Name)
}

// DoctorService ...
type DoctorService struct{}

// Notify ...
func (d *DoctorService) Notify(data interface{}) {
	fmt.Printf("A doctor has been called for %s",
		data.(string))
}

func main() {
	p := NewPerson("Boris")
	ds := &DoctorService{}
	p.Subscribe(ds)

	// let's test it!
	p.CatchACold()
}
