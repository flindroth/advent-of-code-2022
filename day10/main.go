package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/flindroth/advent-of-code-2022/util"
)

func main() {
	lines, err := util.GetPuzzleInput(10)
	if err != nil {
		log.Fatalf("Could not get puzzle input: %v", err.Error())
	}

	x := 1
	pc := 0
	cycle := 1
	inAdd := false
	sig := 0
	render(x, cycle)
	for pc < len(lines) {
		cycle++

		instr := parseInstr(lines[pc])
		//log.Printf("c=%v, pc=%v, instr=%v", cycle, pc, instr)
		switch instr.op {
		case "noop":
			pc++
		case "addx":
			num := util.MustAtoi(instr.oper)
			if inAdd {
				x += num
				inAdd = false
				pc++
			} else {
				inAdd = true
			}
		}

		if cycle == 20 || (cycle-20)%40 == 0 {
			//log.Printf("Signal strength at cycle %v: %v (cycle*x = %v)", cycle, x, cycle*x)
			sig += cycle * x
		}
		render(x, cycle)
	}
	log.Printf("Signal sum: %v", sig)

	/*
		x := 1
		sig := 0
		cycle := 1
		render(x, cycle)
		for _, line := range lines {
			instr := parseInstr(line)
			switch instr.op {
			case "noop":
				cycle++
			case "addx":
				num := util.MustAtoi(instr.oper)
				//log.Printf("cycle=%v, x=%v, adding %v, will be %v", cycle, x, num, x+num)
				cycle++
				render(x, cycle)
				if cycle == 20 || (cycle-20)%40 == 0 {
					//log.Printf("Signal strength at cycle %v (mid-addx): %v (cycle*x = %v)", cycle, x, cycle*x)
					sig += cycle * x
				}
				x += num
				cycle++
			}
			if cycle == 20 || (cycle-20)%40 == 0 {
				//log.Printf("Signal strength at cycle %v: %v (cycle*x = %v)", cycle, x, cycle*x)
				sig += cycle * x
			}
			render(x, cycle)
		}
		log.Printf("Signal sum: %v", sig)
	*/
}

func parseInstr(input string) instr {
	parts := strings.Split(input, " ")
	oper := ""
	if len(parts) == 2 {
		oper = parts[1]
	}
	return instr{parts[0], oper}
}

func render(x, cycle int) {
	crtpos := (cycle - 1) % 40
	if crtpos > x-2 && crtpos < x+2 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
	if crtpos == 39 {
		fmt.Println()
	}
}

type instr struct {
	op   string
	oper string
}
