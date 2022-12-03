package main

import (
	"log"

	"github.com/flindroth/advent-of-code-2022/util"
)

type rucksack string

func main() {
	lines, err := util.GetPuzzleInput(3)
	if err != nil {
		log.Fatalf("Could not fetch puzzle input: %v", err.Error())
	}

	prioSum := 0
	for _, v := range lines {
		r := rucksack(v)
		for _, v := range r.itemTypesInBoth() {
			prioSum += priority(rune(v))
		}
	}

	log.Printf("Prio sum: %v", prioSum)

	groups := splitIntoGroups(lines, 3)
	prioSum = 0
	for _, g := range groups {
		common := commonCharacters(string(commonCharacters(g[0], g[1])), g[2])
		for _, c := range common {
			prioSum += priority(c)
		}

	}

	log.Printf("Prio sum of badges: %v", prioSum)
}

func (r rucksack) firstCompartment() string {
	return string(r[0 : len(r)/2])
}

func (r rucksack) secondCompartment() string {
	return string(r[len(r)/2:])
}

func (r rucksack) itemTypesInBoth() []rune {
	fc := r.firstCompartment()
	sc := r.secondCompartment()

	return commonCharacters(fc, sc)
}

func commonCharacters(str1, str2 string) []rune {
	var itemFlags uint64 = 0
	for _, i := range str1 {
		itemFlags = itemFlags | 1<<int(i-65)
	}

	var itemFlagsSec uint64 = 0
	for _, i := range str2 {
		itemFlagsSec = itemFlagsSec | 1<<int(i-65)
	}

	commonItemFlags := itemFlags & itemFlagsSec

	commonRunes := make([]rune, 0)

	for i := 0; i < 64; i++ {
		bit := commonItemFlags & 1
		if bit == 1 {
			commonRunes = append(commonRunes, rune(int32(i+65)))
		}
		commonItemFlags = commonItemFlags >> 1
	}

	return commonRunes
}

func priority(r rune) int {
	i := int(r)
	if i < 91 { // capital
		return 27 + i - 65
	}
	return 1 + i - 97
}

func splitIntoGroups(list []string, groupSize int) [][]string {
	groups := make([][]string, 0)
	currentGroup := make([]string, 0)
	for _, v := range list {
		currentGroup = append(currentGroup, v)
		if len(currentGroup) == groupSize {
			groups = append(groups, currentGroup)
			currentGroup = make([]string, 0)
		}
	}
	return groups
}
