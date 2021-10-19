package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2

	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

type Round struct {
	Winner         int    `json:"winner"`
	ComputerChoise string `json:"computer_choise"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {

	rand.Seed(time.Now().UnixNano())

	computerValue := rand.Intn(3)
	computerChoise := ""
	roundResult := ""
	winner := 0

	switch computerValue {
	case ROCK:
		computerChoise = "Computer chose ROCK !"
	case PAPER:
		computerChoise = "Computer chose PAPER !"
	case SCISSORS:
		computerChoise = "Computer chose SCISSORS !"
	default:
		computerChoise = "Default message. Computer chose NOTHING !"
	}

	if computerValue == playerValue {
		roundResult = "*** it's a draw ***"
		winner = DRAW
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player Wins!"
		winner = PLAYERWINS
	} else {
		roundResult = "Computer Wins!"
		winner = COMPUTERWINS
	}

	// return output of the function

	var result Round

	result.Winner = winner
	result.ComputerChoise = computerChoise
	result.RoundResult = roundResult

	return result

}
