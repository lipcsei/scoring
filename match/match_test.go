package match

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {

	match := Match{
		Player: []Player{
			{ID: "TP1", Status: "SUB_IN"},
			{ID: "TP1", Status: "SUB_IN"},
			{ID: "TP1", Status: "SUB_IN"},
		},
	}

	fmt.Println(match)

}
