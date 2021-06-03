package main

import "fmt"

// Person ...
type Person struct {
	FirstName, MiddleName, LastName string
}

// Names ...
func (p *Person) Names() []string {
	return []string{p.FirstName, p.MiddleName, p.LastName}
}

// NamesGenerator ...
func (p *Person) NamesGenerator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.FirstName
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()
	return out
}

// PersonNameIterator ...
type PersonNameIterator struct {
	person  *Person
	current int
}

// NewPersonNameIterator ...
func NewPersonNameIterator(person *Person) *PersonNameIterator {
	return &PersonNameIterator{person, -1}
}

// MoveNext ...
func (p *PersonNameIterator) MoveNext() bool {
	p.current++
	return p.current < 3
}

// Value ...
func (p *PersonNameIterator) Value() string {
	switch p.current {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastName
	}
	panic("should not be here")
}

func main() {
	p := Person{"Alexander", "Graham", "Bell"}
	for _, name := range p.Names() {
		fmt.Println(name)
	}

	for name := range p.NamesGenerator() {
		fmt.Println(name)
	}

	for it := NewPersonNameIterator(&p); it.MoveNext(); {
		fmt.Println(it.Value())
	}
}
