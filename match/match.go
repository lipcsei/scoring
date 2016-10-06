package match

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
	ID   string
	Name string
}
