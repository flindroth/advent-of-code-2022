package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/flindroth/advent-of-code-2022/util"
)

func main() {
	lines, err := util.GetPuzzleInput(5)
	if err != nil {
		log.Fatalf("Could not read puzzle input: %v", err.Error())
	}

	crateLines := make([]string, 0)
	instrLines := make([]string, 0)
	isOnInstr := false
	for _, line := range lines {
		if line == "" {
			isOnInstr = true
			continue
		}
		if isOnInstr {
			instrLines = append(instrLines, line)
		} else {
			crateLines = append(crateLines, line)
		}
	}
	log.Printf("Read %v crate lines and %v instr lines", len(crateLines), len(instrLines))

	// Star 1
	stacks := makeStacks(crateLines)

	for _, l := range instrLines {
		amt, from, to := parseInstruction(l)
		for i := 0; i < amt; i++ {
			r := stacks[from-1].Pop()
			stacks[to-1].Push(r)
		}
	}

	final := ""
	for _, s := range stacks {
		final += string(s.Peek())
	}
	log.Println(final)

	// Star 2
	stacks = makeStacks(crateLines)

	for _, l := range instrLines {
		amt, from, to := parseInstruction(l)
		grabbed := []rune{}
		for i := 0; i < amt; i++ {
			r := stacks[from-1].Pop()
			grabbed = append(grabbed, r)
		}
		for i := len(grabbed) - 1; i >= 0; i-- {
			stacks[to-1].Push(grabbed[i])
		}
	}

	final = ""
	for _, s := range stacks {
		final += string(s.Peek())
	}
	log.Println(final)
}

func positionsHavingChars(input string) []int {
	p := []int{}
	for i, c := range input {
		if c != ' ' {
			p = append(p, i)
		}
	}
	return p
}

func parseInstruction(input string) (amt, from, to int) {
	input = strings.ReplaceAll(input, "move ", "")
	input = strings.ReplaceAll(input, " from ", ",")
	input = strings.ReplaceAll(input, " to ", ",")
	parts := strings.Split(input, ",")

	return mustAtoi(parts[0]), mustAtoi(parts[1]), mustAtoi(parts[2])
}

func mustAtoi(a string) int {
	if i, err := strconv.Atoi(a); err == nil {
		return i
	}
	log.Fatalf("Could not convert string \"%v\" to a number", a)
	return -1
}

func makeStacks(crateLines []string) []util.Stack[rune] {
	charPositions := positionsHavingChars(crateLines[len(crateLines)-1])
	stacks := make([]util.Stack[rune], len(charPositions))

	for i := len(crateLines) - 2; i >= 0; i-- {
		for p, j := range charPositions {
			char := rune(crateLines[i][j])

			if char != ' ' {
				stacks[p].Push(char)
			}
		}
	}
	return stacks
}
