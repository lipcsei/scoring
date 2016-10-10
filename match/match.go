package match

import (
	"log"
	"time"

	"github.com/lipcsei/scoring/timemap"
	mgo "gopkg.in/mgo.v2"
)

type Match struct {
	ID      string    `bson:"_id"`
	Kickoff time.Time `bson:"kickoff"`
	Status  string    `bson:"status"`
	Event   []Event   `bson:"event"`
	Player  []Player  `bson:"player"`
}

type Event struct {
	ID       string `bson:"_id"`
	MatchID  string `bson:"matchId"`
	PlayerID string `bson:"playerId"`
	TeamID   string `bson:"teamId"`
	Type     string `bson:"type"`
	Minute   Minute `bson:"minute"`
	Value    Value
}

type Minute struct {
	Minute      int
	MinuteExtra int
}

type Value struct {
	Personal     int `bson:"personal"`
	Team         int `bson:"team"`
	OpponentTeam int `bson:"opponentTeam"`
}

type Player struct {
	ID       string     `bson:"_id"`
	Name     string     `bson:"name"`
	Status   string     `bson:"status"`
	Position string     `bson:"position"`
	Score    ScoreSheet `bson:"score"`
	Value    string     `bson:"value"`
}

type ScoreSheet struct {
	Activity        timemap.TimeMap `bson:"activity"`
	Personal        map[string]int  `bson:"personal"`
	Team            map[string]int  `bson:"team"`
	TeamInLeadBonus timemap.TimeMap `bson:"teamInLeadBonus"`
	DrawBonus       timemap.TimeMap `bson:"drawBonus"`
	CleanSheet      timemap.TimeMap `bson:"cleanSheet"`
	TotalScore      int             `bson:"totalScore"`
	Minutes         int             `bson:"minutes"`
}

func Get(ID string) Match {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/test")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	c := session.DB("test").C("match")

	match := Match{}
	err = c.FindId(ID).One(&match)
	if err != nil {
		log.Fatal(err)
	}

	return match
}
