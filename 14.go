package main

import (
	"fmt"
	"strings"
)

type sandGrid struct {
	bottom int
	sand   *point
	units  int
	grid   [200][1000]byte
}

func day14() {
	lines := mustReadFileLines("14.txt")
	grid := parseGrid(lines)
	for grid.step(grid.bottom) {
	}
	fmt.Println(grid.units)

	grid = parseGrid(lines)
	for r, c := grid.bottom+2, 0; c < len(grid.grid[r]); c++ {
		grid.grid[r][c] = '#'
	}
	for grid.step(grid.bottom + 2) {
	}
	fmt.Println(grid.units)
}

func (g *sandGrid) step(bottom int) (progress bool) {
	if g.sand == nil {
		g.sand = &point{500, 0}
		return true
	}
	for _, next := range []point{
		{g.sand.x, g.sand.y + 1},
		{g.sand.x - 1, g.sand.y + 1},
		{g.sand.x + 1, g.sand.y + 1},
	} {
		if g.grid[next.y][next.x] == '.' {
			g.sand.x = next.x
			g.sand.y = next.y
			return g.sand.y < bottom
		}
	}
	g.grid[g.sand.y][g.sand.x] = 'o'
	g.units++
	sand := g.sand
	g.sand = nil
	return sand.x != 500 || sand.y != 0
}

func parseGrid(lines []string) *sandGrid {
	grid := new(sandGrid)
	for r, row := range grid.grid {
		for c := range row {
			grid.grid[r][c] = '.'
		}
	}
	for _, l := range lines {
		parts := strings.Split(l, " -> ")
		var points []point
		for _, part := range parts {
			x, y, ok := strings.Cut(strings.TrimSpace(part), ",")
			if !ok {
				panic("bad format")
			}
			points = append(points, point{
				x: mustParseInt(x),
				y: mustParseInt(y),
			})
		}
		for i := 0; i < len(points)-1; i++ {
			p0 := points[i]
			p1 := points[i+1]
			if p0.x == p1.x {
				start, end := p0.y, p1.y
				if start > end {
					start, end = end, start
				}
				for y := start; y <= end; y++ {
					grid.grid[y][p0.x] = '#'
					if y > grid.bottom {
						grid.bottom = y
					}
				}
			} else if p0.y == p1.y {
				start, end := p0.x, p1.x
				if start > end {
					start, end = end, start
				}
				for x := start; x <= end; x++ {
					grid.grid[p0.y][x] = '#'
				}
			} else {
				panic("no shared coordinate")
			}
		}
	}
	return grid
}

func (grid *sandGrid) print() {
	for row := 0; row < 14; row++ {
		for col := 480; col < 520; col++ {
			fmt.Printf("%c", grid.grid[row][col])
		}
		fmt.Println()
	}
}
