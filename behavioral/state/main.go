package main

import "fmt"

// Switch ...
type Switch struct {
	State State
}

// NewSwitch ...
func NewSwitch() *Switch {
	return &Switch{NewOffState()}
}

// On ...
func (s *Switch) On() {
	s.State.On(s)
}

// Off ...
func (s *Switch) Off() {
	s.State.Off(s)
}

// State ...
type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

// BaseState ...
type BaseState struct{}

// On ...
func (s *BaseState) On(sw *Switch) {
	fmt.Println("Light is already on")
}

// Off ...
func (s *BaseState) Off(sw *Switch) {
	fmt.Println("Light is already off")
}

// OnState ...
type OnState struct {
	BaseState
}

// NewOnState ...
func NewOnState() *OnState {
	fmt.Println("Light turned on")
	return &OnState{BaseState{}}
}

// Off ...
func (o *OnState) Off(sw *Switch) {
	fmt.Println("Turning light off...")
	sw.State = NewOffState()
}

// OffState ...
type OffState struct {
	BaseState
}

// NewOffState ...
func NewOffState() *OffState {
	fmt.Println("Light turned off")
	return &OffState{BaseState{}}
}

// On ...
func (o *OffState) On(sw *Switch) {
	fmt.Println("Turning light on...")
	sw.State = NewOnState()
}

func main() {
	sw := NewSwitch()
	sw.On()
	sw.Off()
	sw.Off()
}
