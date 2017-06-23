package controllers

import (
	"fmt"
	"net/http"
	"github.com/tspn/tic-tac-toe-AI/tictactoe"
)

func Play(w http.ResponseWriter, r *http.Request) {
	str := r.FormValue("data")
	row := int(str[1] - '0')
	col := int(str[0] - '0')

	fmt.Fprint(w, tictactoe.PlayerPick(row, col))
}
