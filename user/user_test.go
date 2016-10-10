package user

import (
	"fmt"
	"testing"

	"github.com/lipcsei/scoring/match"
	"github.com/lipcsei/scoring/timemap"
)

func TestMain(m *testing.M) {

	match := match.Get("OM2241790")
	fmt.Println(match)

	user := User{
		ID: "TU1",
		ActivePlayers: [4]Player{
			{ID: "OP120695", Activity: timemap.TimeMap{Raw: [4]int{-1, -1, -1, -1}}},
			// {ID: "TP2", Activity: (timemap.TimeMap{-1, -1, -1, -1})},
			// {ID: "TP3", Activity: timemap.TimeMap{}},
			// {ID: "TP4", Activity: timemap.TimeMap{}},
		},
	}
	fmt.Println(user)

}
