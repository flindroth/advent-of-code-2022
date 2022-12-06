package main

import (
	"log"

	"github.com/flindroth/advent-of-code-2022/util"
)

func main() {
	lines, err := util.GetPuzzleInput(6)
	if err != nil {
		log.Fatalf("Could not read puzzle input: %v", err.Error())
	}

	packet := lines[0]

	// Star 1
	for o := 0; o < len(packet)-4; o++ {
		window := packet[o : o+4]
		if unique([]rune(window)) {
			log.Printf("offset: %v, SOP: %v", o, o+4)
			break
		}
	}

	// Star 2
	for o := 0; o < len(packet)-14; o++ {
		window := packet[o : o+14]
		if unique([]rune(window)) {
			log.Printf("offset: %v, SOM: %v", o, o+14)
			break
		}
	}
}

func unique[T comparable](input []T) bool {
	set := util.Set[T]{}
	for _, i := range input {
		set.Add(i)
	}
	return len(set) == len(input)
}
