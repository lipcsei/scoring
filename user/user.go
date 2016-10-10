package user

import "github.com/lipcsei/scoring/timemap"

type User struct {
	ID            string
	ActivePlayers [4]Player
	UsedPlayers   []Player
}

type Player struct {
	ID       string
	Activity timemap.TimeMap `bson:"activity"`
}
