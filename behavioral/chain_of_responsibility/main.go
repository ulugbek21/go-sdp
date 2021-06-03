package main

import "fmt"

// Creature ...
type Creature struct {
	Name            string
	Attack, Defense int
}

// String ...
func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)",
		c.Name, c.Attack, c.Defense)
}

// NewCreature ...
func NewCreature(name string, attack int, defense int) *Creature {
	return &Creature{Name: name, Attack: attack, Defense: defense}
}

// Modifier ...
type Modifier interface {
	Add(m Modifier)
	Handle()
}

// CreatureModifier ...
type CreatureModifier struct {
	creature *Creature
	next     Modifier // singly linked list
}

// Add ...
func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

// Handle ...
func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

// NewCreatureModifier ...
func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{creature: creature}
}

// DoubleAttackModifier ...
type DoubleAttackModifier struct {
	CreatureModifier
}

// NewDoubleAttackModifier ...
func NewDoubleAttackModifier(
	c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{
		creature: c}}
}

// IncreasedDefenseModifier ...
type IncreasedDefenseModifier struct {
	CreatureModifier
}

// NewIncreasedDefenseModifier ...
func NewIncreasedDefenseModifier(
	c *Creature) *IncreasedDefenseModifier {
	return &IncreasedDefenseModifier{CreatureModifier{
		creature: c}}
}

// Handle ...
func (i *IncreasedDefenseModifier) Handle() {
	if i.creature.Attack <= 2 {
		fmt.Println("Increasing",
			i.creature.Name, "\b's defense")
		i.creature.Defense++
	}
	i.CreatureModifier.Handle()
}

// Handle ...
func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling", d.creature.Name,
		"attack...")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()
}

// NoBonusesModifier ...
type NoBonusesModifier struct {
	CreatureModifier
}

// NewNoBonusesModifier ...
func NewNoBonusesModifier(
	c *Creature) *NoBonusesModifier {
	return &NoBonusesModifier{CreatureModifier{
		creature: c}}
}

// Handle ...
func (n *NoBonusesModifier) Handle() {
	// nothing here!
}

func main() {
	goblin := NewCreature("Goblin", 1, 1)
	fmt.Println(goblin)

	root := NewCreatureModifier(goblin)

	//root.Add(NewNoBonusesModifier(goblin))

	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewIncreasedDefenseModifier(goblin))
	root.Add(NewDoubleAttackModifier(goblin))

	// eventually process the entire chain
	root.Handle()
	fmt.Println(goblin)
}
