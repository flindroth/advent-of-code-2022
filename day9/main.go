package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/flindroth/advent-of-code-2022/util"
)

func main() {
	lines, err := util.GetPuzzleInput(9)
	if err != nil {
		log.Fatalf("Could not get puzzle input: %v", err.Error())
	}

	hx := 0
	hy := 0
	tx := 0
	ty := 0
	visited := make(util.Set[string])
	for _, dir := range getMovements(lines) {
		log.Printf("Motion: %v", string(dir))
		switch dir {
		case 'U':
			hy++
			if !isTouching(hx, hy, tx, ty) {
				tx = hx
				ty = hy - 1
			}

		case 'D':
			hy--
			if !isTouching(hx, hy, tx, ty) {
				tx = hx
				ty = hy + 1
			}

		case 'L':
			hx--
			if !isTouching(hx, hy, tx, ty) {
				ty = hy
				tx = hx + 1
			}

		case 'R':
			hx++
			if !isTouching(hx, hy, tx, ty) {
				ty = hy
				tx = hx - 1
			}
		}
		visited.Add(fmt.Sprintf("%v,%v", tx, ty))
		log.Printf("After: Head=%v,%v, Tail=%v,%v", hx, hy, tx, ty)
	}
	log.Printf("Tail visited %v positions", len(visited))
}

func parseMotion(line string) (rune, int) {
	parts := strings.Split(line, " ")
	return rune(parts[0][0]), mustAtoi(parts[1])
}

func getMovements(lines []string) []rune {
	m := []rune{}
	for _, l := range lines {
		dir, amt := parseMotion(l)
		for i := 0; i < amt; i++ {
			m = append(m, dir)
		}
	}
	return m
}

func isTouching(hx, hy, tx, ty int) bool {
	dx := hx - tx
	dy := hy - ty
	return dx <= 1 && dx >= -1 && dy <= 1 && dy >= -1

}

type segment struct {
	x int
	y int
}

func mustAtoi(number string) int {
	i, err := strconv.Atoi(number)
	if err != nil {
		log.Fatalf("Could not convert string \"%v\" to number", number)
	}
	return i
}
