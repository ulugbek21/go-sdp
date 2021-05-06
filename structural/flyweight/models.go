package flyweight

import "time"

// Team ...
type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

const (
	// TeamA ...
	TeamA = iota
	// TeamB ...
	TeamB
)

// Player ...
type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

// HistoricalData ...
type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

// Match ...
type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}
