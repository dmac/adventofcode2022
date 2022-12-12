package main

import (
	"fmt"
	"math"
)

func day12() {
	var grid [][]byte
	lines := mustReadFileLines("12.txt")
	for _, l := range lines {
		grid = append(grid, []byte(l))
	}
	var start, end point
	for y, row := range grid {
		for x, e := range row {
			if e == 'S' {
				start = point{x: x, y: y}
				grid[y][x] = 'a'
			}
			if e == 'E' {
				end = point{x: x, y: y}
				grid[y][x] = 'z'
			}
		}
	}

	part1 := findGoal(grid, start, end)
	fmt.Println(part1)

	min := math.MaxInt
	for y, row := range grid {
		for x, e := range row {
			if e != 'a' && e != 'S' {
				continue
			}
			p := point{x, y}
			steps := findGoal(grid, p, end)
			if steps < min {
				min = steps
			}
		}
	}
	fmt.Println(min)
}

func findGoal(grid [][]byte, start, end point) int {
	type search struct {
		p    point
		step int
	}
	var curr search
	frontier := []search{{p: start}}
	visited := make(map[point]struct{})
	visited[start] = struct{}{}
	checkPoint := func(p point) {
		if _, ok := visited[p]; ok {
			return
		}
		if grid[p.y][p.x] > grid[curr.p.y][curr.p.x]+1 {
			return
		}
		next := search{
			p:    p,
			step: curr.step + 1,
		}
		visited[p] = struct{}{}
		frontier = append(frontier, next)
	}
	for len(frontier) > 0 {
		curr, frontier = frontier[0], frontier[1:]
		if curr.p == end {
			return curr.step
		}
		if curr.p.y > 0 {
			p := point{x: curr.p.x, y: curr.p.y - 1}
			checkPoint(p)
		}
		if curr.p.y < len(grid)-1 {
			p := point{x: curr.p.x, y: curr.p.y + 1}
			checkPoint(p)
		}
		if curr.p.x > 0 {
			p := point{x: curr.p.x - 1, y: curr.p.y}
			checkPoint(p)
		}
		if curr.p.x < len(grid[0])-1 {
			p := point{x: curr.p.x + 1, y: curr.p.y}
			checkPoint(p)
		}
	}
	return math.MaxInt
}

func printGrid(grid [][]byte, pos point) {
	for y, row := range grid {
		for x, e := range row {
			p := point{x: x, y: y}
			if p == pos {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c", e)
			}
		}
		fmt.Println()
	}
}
