package util

import (
	"log"
	"strconv"
)

func MustAtoi(input string) int {
	if i, err := strconv.Atoi(input); err == nil {
		return i
	}
	log.Fatalf("Could not parse string \"input\" as an integer")
	return -1
}
