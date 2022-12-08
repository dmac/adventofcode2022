package main

import "fmt"

func day8() {
	lines := mustReadFileLines("8.txt")
	numRows := len(lines)
	numCols := len(lines[0])
	grid := make([]int, numRows*numCols)
	for row, line := range lines {
		for col, b := range line {
			grid[row*numCols+col] = int(b) - '0'
		}
	}
	visible := make([]bool, len(grid))
	for row := 0; row < numRows; row++ {
		leftMax := -1
		for col := 0; col < numCols; col++ {
			h := grid[row*numCols+col]
			if h > leftMax {
				visible[row*numCols+col] = true
				leftMax = h
			}
		}
	}
	for row := 0; row < numRows; row++ {
		rightMax := -1
		for col := numCols - 1; col >= 0; col-- {
			h := grid[row*numCols+col]
			if h > rightMax {
				visible[row*numCols+col] = true
				rightMax = h
			}
		}
	}
	for col := 0; col < numCols; col++ {
		topMax := -1
		for row := 0; row < numRows; row++ {
			h := grid[row*numCols+col]
			if h > topMax {
				visible[row*numCols+col] = true
				topMax = h
			}
		}
	}
	for col := 0; col < numCols; col++ {
		bottomMax := -1
		for row := numRows - 1; row >= 0; row-- {
			h := grid[row*numCols+col]
			if h > bottomMax {
				visible[row*numCols+col] = true
				bottomMax = h
			}
		}
	}

	part1 := 0
	for _, vis := range visible {
		if vis {
			part1++
		}
	}
	fmt.Println(part1)

	maxScenic := 0
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			scenic := computeScenic(grid, numCols, numRows, row, col)
			if scenic > maxScenic {
				maxScenic = scenic
			}
		}
	}
	fmt.Println(maxScenic)
}

func computeScenic(grid []int, numCols, numRows, startRow, startCol int) int {
	h := grid[startRow*numCols+startCol]
	lookUp := 0
	for row := startRow - 1; row >= 0; row-- {
		lookUp++
		hh := grid[row*numCols+startCol]
		if hh >= h {
			break
		}
	}
	lookDown := 0
	for row := startRow + 1; row < numRows; row++ {
		lookDown++
		hh := grid[row*numCols+startCol]
		if hh >= h {
			break
		}
	}
	lookLeft := 0
	for col := startCol - 1; col >= 0; col-- {
		lookLeft++
		hh := grid[startRow*numCols+col]
		if hh >= h {
			break
		}
	}
	lookRight := 0
	for col := startCol + 1; col < numCols; col++ {
		lookRight++
		hh := grid[startRow*numCols+col]
		if hh >= h {
			break
		}
	}
	return lookUp * lookDown * lookLeft * lookRight
}
