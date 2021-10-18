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

func PlayRound(playerValue int) (int, string, string) {
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

	return winner, computerChoise, roundResult
}
