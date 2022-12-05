package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func day5() {
	lines := mustReadFileLines("5.txt")
	var numStacks int
	var stackBase int
	for i, l := range lines {
		if strings.TrimSpace(l) == "" {
			stackBase = i - 2
			prev := strings.TrimSpace(lines[i-1])
			var err error
			numStacks, err = strconv.Atoi(string(prev[len(prev)-1]))
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	}
	if numStacks == 0 {
		panic("unknown stacks")
	}
	stacks1 := make([][]byte, numStacks)
	stacks2 := make([][]byte, numStacks)
	for i := 0; i < numStacks; i++ {
		for j := stackBase; j >= 0; j-- {
			l := lines[j]
			if i*4+1 >= len(l) {
				continue
			}
			c := l[i*4+1]
			if c == ' ' {
				continue
			}
			stacks1[i] = append(stacks1[i], c)
			stacks2[i] = append(stacks2[i], c)
		}
	}

	execute(stacks1, lines, stackBase+3, false)
	printTop(stacks1)
	execute(stacks2, lines, stackBase+3, true)
	printTop(stacks2)
}

func execute(stacks [][]byte, lines []string, start int, ordered bool) {
	for i := start; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		n := mustParseInt(parts[1])
		src := mustParseInt(parts[3])
		dst := mustParseInt(parts[5])
		if ordered {
			moving := stacks[src-1][len(stacks[src-1])-n:]
			stacks[src-1] = stacks[src-1][:len(stacks[src-1])-n]
			stacks[dst-1] = append(stacks[dst-1], moving...)
		} else {
			for j := 0; j < n; j++ {
				c := stacks[src-1][len(stacks[src-1])-1]
				stacks[src-1] = stacks[src-1][:len(stacks[src-1])-1]
				stacks[dst-1] = append(stacks[dst-1], c)
			}
		}

	}
}

func printTop(stacks [][]byte) {
	for _, s := range stacks {
		fmt.Printf("%c", s[len(s)-1])
	}
	fmt.Println()
}
