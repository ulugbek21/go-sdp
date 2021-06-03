package main

import "fmt"

// Person ...
type Person struct {
	Name    string
	Room    *ChatRoom
	chatLog []string
}

// NewPerson ...
func NewPerson(name string) *Person {
	return &Person{Name: name}
}

// Receive ...
func (p *Person) Receive(sender, message string) {
	s := fmt.Sprintf("%s: '%s'", sender, message)
	fmt.Printf("[%s's chat session] %s\n", p.Name, s)
	p.chatLog = append(p.chatLog, s)
}

// Say ...
func (p *Person) Say(message string) {
	p.Room.Broadcast(p.Name, message)
}

// PrivateMessage ...
func (p *Person) PrivateMessage(who, message string) {
	p.Room.Message(p.Name, who, message)
}

// ChatRoom ...
type ChatRoom struct {
	people []*Person
}

// Broadcast ...
func (c *ChatRoom) Broadcast(source, message string) {
	for _, p := range c.people {
		if p.Name != source {
			p.Receive(source, message)
		}
	}
}

// Join ...
func (c *ChatRoom) Join(p *Person) {
	joinMsg := p.Name + " joins the chat"
	c.Broadcast("Room", joinMsg)

	p.Room = c
	c.people = append(c.people, p)
}

// Message ...
func (c *ChatRoom) Message(src, dst, msg string) {
	for _, p := range c.people {
		if p.Name == dst {
			p.Receive(src, msg)
		}
	}
}

func main() {
	room := ChatRoom{}

	john := NewPerson("John")
	jane := NewPerson("Jane")

	room.Join(john)
	room.Join(jane)

	john.Say("hi room")
	jane.Say("oh, hey john")

	simon := NewPerson("Simon")
	room.Join(simon)
	simon.Say("hi everyone!")

	jane.PrivateMessage("Simon", "glad you could join us!")
}
