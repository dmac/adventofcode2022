package main

import (
	"fmt"
	"strings"
)

type point struct {
	x int
	y int
}

func day9() {
	ropeSimulation(2)
	ropeSimulation(10)
}

func ropeSimulation(knots int) {
	rope := make([]point, knots)
	tailVisited := make(map[point]struct{})
	tailVisited[rope[len(rope)-1]] = struct{}{}

	lines := mustReadFileLines("9.txt")
	for _, l := range lines {
		dir, ns, _ := strings.Cut(l, " ")
		n := mustParseInt(ns)
		for i := 0; i < n; i++ {
			head := &rope[0]
			switch dir {
			case "U":
				head.y++
			case "D":
				head.y--
			case "L":
				head.x--
			case "R":
				head.x++
			}
			for j := 0; j < len(rope)-1; j++ {
				head := &rope[j]
				tail := &rope[j+1]
				stepTail(head, tail)
				if j == len(rope)-2 {
					tailVisited[*tail] = struct{}{}
				}
			}
		}
	}
	fmt.Println(len(tailVisited))
}

func stepTail(head, tail *point) {
	switch {
	case head.x == tail.x && head.y == tail.y+2:
		tail.y++
		return
	case head.x == tail.x && head.y == tail.y-2:
		tail.y--
		return
	case head.y == tail.y && head.x == tail.x+2:
		tail.x++
		return
	case head.y == tail.y && head.x == tail.x-2:
		tail.x--
		return
	}
	dx := head.x - tail.x
	dy := head.y - tail.y
	if abs(dx) < 2 && abs(dy) < 2 {
		return
	}
	if dx > 0 {
		tail.x++
	} else if dx < 0 {
		tail.x--
	}
	if dy > 0 {
		tail.y++
	} else if dy < 0 {
		tail.y--
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
