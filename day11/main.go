package main

import (
	"log"
	"sort"
	"strings"

	"github.com/flindroth/advent-of-code-2022/util"
)

var (
	monkeys  []monkey
	inspects []int
)

func main() {
	lines, err := util.GetPuzzleInput(11)
	if err != nil {
		log.Fatalf("Could not get puzzle input: %v", err.Error())
	}

	monkeys = parseMonkeys(lines)
	inspects = make([]int, len(monkeys))
	log.Printf("Parsed %v monkeys", len(monkeys))

	for r := 0; r < 20; r++ {
		for i := range monkeys {
			performTurn(i, -1)
		}
	}

	sort.Ints(inspects)
	log.Printf("Monkey multiplication: %v", inspects[len(inspects)-1]*inspects[len(inspects)-2])

	monkeys = parseMonkeys(lines)
	gcd := calculateGCD(monkeys)
	inspects = make([]int, len(monkeys))
	for r := 0; r < 10000; r++ {
		for i := range monkeys {
			performTurn(i, gcd)
		}
	}

	sort.Ints(inspects)
	log.Printf("Monkey multiplication: %v", inspects[len(inspects)-1]*inspects[len(inspects)-2])
}

func parseMonkeys(input []string) []monkey {
	monkeys := []monkey{}
	lineChunks := util.ChunkSlice(input, 7)
	for _, c := range lineChunks {
		monkeys = append(monkeys, parseMonkey(c))
	}
	return monkeys
}

func parseMonkey(input []string) monkey {
	itemsLine := strings.TrimPrefix(input[1], "  Starting items: ")
	itemStrings := strings.Split(itemsLine, ", ")
	items := []int{}
	for _, s := range itemStrings {
		items = append(items, util.MustAtoi(s))
	}
	opString := strings.TrimPrefix(input[2], "  Operation: ")
	opTokens := strings.Split(opString, " ")
	leftOper := opTokens[2]
	op := rune(opTokens[3][0])
	rightOper := opTokens[4]
	div := util.MustAtoi(strings.TrimPrefix(input[3], "  Test: divisible by "))
	testPassMonkey := util.MustAtoi(strings.TrimPrefix(input[4], "    If true: throw to monkey "))
	testFailMonkey := util.MustAtoi(strings.TrimPrefix(input[5], "    If false: throw to monkey "))

	return monkey{
		items,
		leftOper,
		op,
		rightOper,
		div,
		testPassMonkey,
		testFailMonkey,
	}

}

func performTurn(mnum int, gcd int) {
	m := monkeys[mnum]
	for _, i := range m.items {
		if gcd != -1 {
			i = i % gcd
		}
		//log.Printf("Monkey %v inspects item %v", mnum, i)
		val := calc(i, m.leftOper, m.op, m.rightOper, gcd)
		if val%m.div == 0 {
			monkeys[m.testPassMonkey].items = append(monkeys[m.testPassMonkey].items, val)
		} else {
			monkeys[m.testFailMonkey].items = append(monkeys[m.testFailMonkey].items, val)
		}

	}
	inspects[mnum] = inspects[mnum] + len(m.items)
	monkeys[mnum].items = []int{}
}

func calc(val int, leftOper string, op rune, rightOper string, gcd int) int {
	left := parseOper(val, leftOper)
	right := parseOper(val, rightOper)
	res := 0
	switch op {
	case '+':
		res = left + right
	case '*':
		res = left * right
	}
	if gcd != -1 {
		return res % gcd
	}
	return res / 3
}

func parseOper(old int, oper string) int {
	i := 0
	if oper == "old" {
		i = old
	} else {
		i = util.MustAtoi(oper)
	}
	return i
}

func calculateGCD(monkeys []monkey) int {
	gcd := 1
	for _, m := range monkeys {
		gcd *= m.div
	}
	return gcd
}

type monkey struct {
	items          []int
	leftOper       string
	op             rune
	rightOper      string
	div            int
	testPassMonkey int
	testFailMonkey int
}
