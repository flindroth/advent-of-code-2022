package util

import (
	"log"
	"strconv"
)

func MustAtoi(input string) int {
	if i, err := strconv.Atoi(input); err == nil {
		return i
	}
	log.Fatalf("Could not parse string \"%v\" as an integer", input)
	return -1
}
