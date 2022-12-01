package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

const (
	filePath = "input.txt"
)

var (
	maxCal    = 0
	calCounts = make([]int, 0)
)

func main() {
	curCal := 0
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elfDone(curCal)
			curCal = 0
		}
		if i, err := strconv.Atoi(line); err == nil {
			curCal += i
		}
	}
	elfDone(curCal)
	log.Printf("Max cal: %v\n", maxCal)

	sort.Ints(calCounts)
	threeTop := calCounts[len(calCounts)-3:]

	log.Printf("Sum of elves with most cals: %v", sum(threeTop))
}

func elfDone(elfCal int) {
	if elfCal > maxCal {
		maxCal = elfCal
	}
	calCounts = append(calCounts, elfCal)
}

func sum(ints []int) int {
	s := 0
	for _, v := range ints {
		s += v
	}
	return s
}
