package match

import "github.com/lipcsei/scoring/timemap"

type Match struct {
	ID string

	Status string
	Event  []Event
	Player []Player
}

type Event struct {
	ID       string
	MatchID  string
	PlayerID string
	TeamID   string
	Type     string
	Minute   Minute
	Value    Value
}

type Minute struct {
	Minute      int
	MinuteExtra int
}

type Value struct {
	Personal     int
	Team         int
	opponentTeam int
}

type Player struct {
	ID     string
	Name   string
	Status string
}

type ScoreSheet struct {
	Activity        timemap.TimeMap
	Personal        map[string]int
	Team            map[string]int
	TeamInLeadBonus timemap.TimeMap
	DrawBonus       timemap.TimeMap
	CleanSheet      timemap.TimeMap
	TotalScore      int
	Minutes         int
}
