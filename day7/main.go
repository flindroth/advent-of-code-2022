package main

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/flindroth/advent-of-code-2022/util"
)

const (
	fsSize   = 70000000
	updNeeds = 30000000
)

var (
	curLine = 0
	dirs    = make(map[string]int)
	lines   []string
)

func main() {
	if l, err := util.GetPuzzleInput(7); err == nil {
		lines = l
	} else {
		log.Fatalf("Could not read puzzle input: %v", err.Error())
	}

	curLine = 2
	processDir("/")

	// Star 1
	sum := 0
	for _, v := range dirs {
		if v <= 100000 {
			sum += v
		}
	}

	log.Printf("SUM=%v", sum)

	// Star 2
	usedSpace := dirs["/"]
	freeSpace := fsSize - usedSpace
	needToFree := updNeeds - freeSpace

	sizes := []int{}
	for _, v := range dirs {
		sizes = append(sizes, v)
	}

	sort.Ints(sizes)
	for _, v := range sizes {
		if v >= needToFree {
			log.Printf("Smallest delete=%v", v)
			break
		}
	}

}

func processDir(p string) {
	fileSize := 0
	for curLine < len(lines) {
		line := lines[curLine]
		if lines[curLine] == "$ cd .." {
			break
		}
		if strings.HasPrefix(line, "dir") {
			curLine++
			continue
		}
		if strings.HasPrefix(line, "$ cd ") {
			dir := strings.TrimPrefix(line, "$ cd ")
			curLine += 2
			processDir(p + "/" + dir)
			fileSize += dirs[p+"/"+dir]
		} else {
			fileSize += mustAtoi(strings.Split(line, " ")[0])
		}
		curLine++
	}
	dirs[p] = fileSize
}

func mustAtoi(number string) int {
	i, err := strconv.Atoi(number)
	if err != nil {
		log.Fatalf("Could not convert string \"%v\" to number", number)
	}
	return i
}
