package main

import (
	"bufio"
	"log"
	"os"
)

type round struct {
	opponent uint8
	you      uint8
}

const (
	filePath = "input.txt"
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
	LOSE     = 0
	DRAW     = 1
	WIN      = 2
)

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	rounds := make([]round, 0)

	for scanner.Scan() {
		line := scanner.Text()
		rounds = append(rounds, parseLine(line))
		//score += scoreRound(currentPlay)
	}

	score := 0
	for _, r := range rounds {
		score += scoreRound(r)
	}

	log.Printf("Score: %v", score)

	score = 0

	strategise(rounds)

	for _, r := range rounds {
		score += scoreRound(r)
	}

	log.Printf("Stragegy score: %v", score)

}

func sum(ints []int) int {
	s := 0
	for _, v := range ints {
		s += v
	}
	return s
}

func parseLine(line string) round {
	firstChar := line[0]
	secondChar := line[2]

	var opponent uint8
	var you uint8
	switch firstChar {
	case 'A':
		opponent = ROCK
	case 'B':
		opponent = PAPER
	case 'C':
		opponent = SCISSORS
	}
	switch secondChar {
	case 'X':
		you = ROCK
	case 'Y':
		you = PAPER
	case 'Z':
		you = SCISSORS
	}

	return round{opponent: opponent, you: you}
}

func scoreRound(r round) int {
	var shapeScore uint8
	switch r.you {
	case ROCK:
		shapeScore = 1
	case PAPER:
		shapeScore = 2
	case SCISSORS:
		shapeScore = 3
	}

	if r.you == r.opponent {
		return int(shapeScore + 3)
	}

	var roundScore uint8 = 0

	switch r.opponent {
	case ROCK:
		if r.you == PAPER {
			roundScore = 6
		}
	case PAPER:
		if r.you == SCISSORS {
			roundScore = 6
		}
	case SCISSORS:
		if r.you == ROCK {
			roundScore = 6
		}
	}

	return int(shapeScore + roundScore)
}

func strategise(rounds []round) {
	for i, r := range rounds {
		var newMove uint8
		switch r.you {
		case DRAW:
			newMove = r.opponent
		case LOSE:
			switch r.opponent {
			case ROCK:
				newMove = SCISSORS
			case PAPER:
				newMove = ROCK
			case SCISSORS:
				newMove = PAPER
			}
		case WIN:
			switch r.opponent {
			case ROCK:
				newMove = PAPER
			case PAPER:
				newMove = SCISSORS
			case SCISSORS:
				newMove = ROCK
			}
		}
		rounds[i] = round{opponent: r.opponent, you: newMove}
	}
}
