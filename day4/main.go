package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/flindroth/advent-of-code-2022/util"
)

func main() {
	lines, err := util.GetPuzzleInput(4)
	if err != nil {
		log.Fatalf("Could not get puzzle input: %v", err.Error())
	}

	pairs := make([][]util.Set[int], 0)
	for _, line := range lines {
		elf1Range, elf2Range := parsePair(line)
		pairs = append(pairs, []util.Set[int]{elf1Range, elf2Range})
	}

	// Star 1
	numOfContainsAll := 0
	for _, pair := range pairs {
		elf1Range := pair[0]
		elf2Range := pair[1]

		intersect := elf1Range.Intersect(elf2Range)

		if intersect.Equals(elf1Range) || intersect.Equals(elf2Range) {
			numOfContainsAll++
		}
	}
	log.Printf("Number of \"contains all\": %v", numOfContainsAll)

	// Star 2
	overlaps := 0
	for _, pair := range pairs {
		elf1Range := pair[0]
		elf2Range := pair[1]

		intersect := elf1Range.Intersect(elf2Range)

		if len(intersect) > 0 {
			overlaps++
		}
	}
	log.Printf("Number of overlaps: %v", overlaps)
}

func parsePair(pair string) (util.Set[int], util.Set[int]) {
	sets := strings.Split(pair, ",")

	elf1 := strings.Split(sets[0], "-")
	elf2 := strings.Split(sets[1], "-")

	elf1Range := util.IntSet(mustAtoi(elf1[0]), mustAtoi(elf1[1]))
	elf2Range := util.IntSet(mustAtoi(elf2[0]), mustAtoi(elf2[1]))

	return elf1Range, elf2Range
}

func mustAtoi(a string) int {
	if i, err := strconv.Atoi(a); err == nil {
		return i
	}
	log.Fatalf("Could not convert string \"%v\" to a number", a)
	return -1
}
