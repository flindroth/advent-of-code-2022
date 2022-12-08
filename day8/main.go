package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/flindroth/advent-of-code-2022/util"
)

func main() {
	lines, err := util.GetPuzzleInput(8)
	if err != nil {
		log.Fatalf("Could not get puzzle input: %v", err.Error())
	}

	trees := [][]int{}

	for _, line := range lines {
		trees = append(trees, parseTreeLine(line))
	}

	// Star 1
	maxLR, maxRL, maxUD, maxDU := makeMaxMaps(trees)
	visible := 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			if i == 0 || i == len(trees)-1 {
				visible++
				continue
			} else if j == 0 || j == len(trees[0])-1 {
				visible++
				continue
			}

			tree := trees[i][j]

			if tree > maxLR[i][j-1] ||
				tree > maxRL[i][j+1] ||
				tree > maxUD[i-1][j] ||
				tree > maxDU[i+1][j] {

				visible++
			}
		}
	}
	log.Printf("Visible: %v", visible)

	// Star 2
	maxScenic := 0
	for i := 0; i < len(trees); i++ {
		maxRowScenic := 0
		for j := 0; j < len(trees[i]); j++ {
			leftOf := trees[i][0:j]
			rightOf := trees[i][j+1:]
			column := getColumn(trees, j)
			above := column[0:i]
			below := column[i+1:]

			tree := trees[i][j]

			scenicLeft := calcScenic(rev(leftOf), tree)
			scenicRight := calcScenic(rightOf, tree)
			scenicUp := calcScenic(rev(above), tree)
			scenicDown := calcScenic(below, tree)

			score := scenicLeft * scenicRight * scenicUp * scenicDown
			if score > maxRowScenic {
				maxRowScenic = score
			}
		}
		if maxRowScenic > maxScenic {
			maxScenic = maxRowScenic
		}
	}
	log.Printf("Max scenic: %v", maxScenic)

}

func parseTreeLine(line string) []int {
	ints := []int{}
	for _, c := range line {
		ints = append(ints, mustAtoi(string(c)))
	}
	return ints
}

func makeMaxMaps(trees [][]int) ([][]int, [][]int, [][]int, [][]int) {
	rows := len(trees)
	cols := len(trees[0])
	maxLR := mkIntMatrix(cols, rows)
	maxRL := mkIntMatrix(cols, rows)
	maxUD := mkIntMatrix(cols, rows)
	maxDU := mkIntMatrix(cols, rows)

	for r := 0; r < rows; r++ {
		for c, n := range runningMax(trees[r]) {
			maxLR[r][c] = n
		}
	}

	for r := 0; r < rows; r++ {
		for c, n := range rev(runningMax(rev(trees[r]))) {
			maxRL[r][c] = n
		}
	}

	for c := 0; c < cols; c++ {
		for r, n := range runningMax(getColumn(trees, c)) {
			maxUD[r][c] = n
		}
	}

	for c := 0; c < cols; c++ {
		for r, n := range rev(runningMax(rev(getColumn(trees, c)))) {
			maxDU[r][c] = n
		}
	}

	return maxLR, maxRL, maxUD, maxDU
}

func runningMax(input []int) []int {
	max := input[0]
	output := []int{input[0]}
	for i := 1; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}
		output = append(output, max)
	}
	return output
}

func getColumn(trees [][]int, col int) []int {
	t := []int{}
	for i := 0; i < len(trees); i++ {
		t = append(t, trees[i][col])
	}
	return t
}

func mkIntMatrix(sx, sy int) [][]int {
	m := make([][]int, sy)
	for r := 0; r < sy; r++ {
		m[r] = make([]int, sx)
	}
	return m
}

func calcScenic(treesInDir []int, thisTree int) int {
	score := 0
	for t := 0; t < len(treesInDir); t++ {
		score += 1
		if treesInDir[t] >= thisTree {
			break
		}
	}
	return score
}

func max(input []int) int {
	max := 0
	for _, v := range input {
		if v > max {
			max = v
		}
	}
	return max
}

func mustAtoi(number string) int {
	i, err := strconv.Atoi(number)
	if err != nil {
		log.Fatalf("Could not convert string \"%v\" to number", number)
	}
	return i
}

func numberLine(input []int) string {
	s := ""
	for _, i := range input {
		s += fmt.Sprintf("%v", i)
	}
	return s
}

func rev[T any](input []T) []T {
	rev := []T{}
	for i := len(input) - 1; i >= 0; i-- {
		rev = append(rev, input[i])
	}
	return rev
}
