package main

import (
	"fmt"
	"log"
)

const (
	rock = iota
	paper
	scissors
)

func day2() {
	points1 := 0
	points2 := 0
	scanner := makeFileScanner("2.txt")
	for scanner.Scan() {
		var opp, me int
		b := scanner.Bytes()
		switch b[0] {
		case 'A':
			opp = rock
		case 'B':
			opp = paper
		case 'C':
			opp = scissors
		default:
			panic("unknown shape")
		}

		// part 1
		switch b[2] {
		case 'X':
			me = rock
			points1 += 1
		case 'Y':
			me = paper
			points1 += 2
		case 'Z':
			me = scissors
			points1 += 3
		default:
			panic("unknown shape")
		}
		if me == opp {
			points1 += 3
		} else if (opp+1)%3 == me {
			points1 += 6
		}

		// part 2
		switch b[2] {
		case 'X':
			me = (opp + 2) % 3
		case 'Y':
			me = opp
			points2 += 3
		case 'Z':
			me = (opp + 1) % 3
			points2 += 6
		default:
			panic("unknown shape")
		}
		switch me {
		case rock:
			points2 += 1
		case paper:
			points2 += 2
		case scissors:
			points2 += 3
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(points1)
	fmt.Println(points2)
}
