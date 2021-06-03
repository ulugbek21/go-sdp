package main

import "fmt"

// Game ...
type Game interface {
	Start()
	HaveWinner() bool
	TakeTurn()
	WinningPlayer() int
}

// PlayGame ...
func PlayGame(g Game) {
	g.Start()
	for !g.HaveWinner() {
		g.TakeTurn()
	}
	fmt.Printf("Player %d wins.\n", g.WinningPlayer())
}

type chess struct {
	turn, maxTurns, currentPlayer int
}

// NameGameOfChess ...
func NewGameOfChess() Game {
	return &chess{1, 10, 0}
}

// Start ...
func (c *chess) Start() {
	fmt.Println("Starting a game of chess.")
}

// HaveWinner ...
func (c *chess) HaveWinner() bool {
	return c.turn == c.maxTurns
}

// TakeTurn ...
func (c *chess) TakeTurn() {
	c.turn++
	fmt.Printf("Turn %d taken by player %d\n",
		c.turn, c.currentPlayer)
	c.currentPlayer = (c.currentPlayer + 1) % 2
}

// WinningPlayer ...
func (c *chess) WinningPlayer() int {
	return c.currentPlayer
}

func main() {
	chess := NewGameOfChess()
	PlayGame(chess)
}
