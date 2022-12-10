package main

import (
	"fmt"
	"strings"
)

type CPU struct {
	X     int
	cycle int

	signalSums int
}

func day10() {
	lines := mustReadFileLines("10.txt")
	var cpu CPU
	cpu.X = 1
	for _, l := range lines {
		inst, ns, _ := strings.Cut(l, " ")
		switch inst {
		case "noop":
			cpu.noop()
		case "addx":
			n := mustParseInt(ns)
			cpu.addx(n)
		}
	}
	fmt.Println()
	fmt.Println(cpu.signalSums)
}

func (c *CPU) noop() {
	c.cycle++
	c.drawPixel()
	if c.signalCycle() {
		c.signalSums += c.X * c.cycle
	}
}

func (c *CPU) addx(n int) {
	c.cycle++
	c.drawPixel()
	if c.signalCycle() {
		c.signalSums += c.X * c.cycle
	}
	c.cycle++
	c.drawPixel()
	if c.signalCycle() {
		c.signalSums += c.X * c.cycle
	}
	c.X += n
}

func (c *CPU) signalCycle() bool {
	return c.cycle == 20 || (c.cycle-20)%40 == 0
}

func (c *CPU) drawPixel() {
	col := (c.cycle % 40) - 1
	if c.cycle > 1 && col == 0 {
		fmt.Println()
	}
	if col == c.X-1 || col == c.X || col == c.X+1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
}
