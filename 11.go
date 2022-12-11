package main

import (
	"fmt"
	"sort"
	"strings"
)

type monkey struct {
	items         []int
	update        func(int) int
	testDivisible int
	testTrue      int
	testFalse     int
	count         int
}

func day11() {
	lines := mustReadFileLines("11small.txt")
	monkeys := parseMonkeys(lines)
	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			m.inspectItems(monkeys)
		}
	}
	for _, m := range monkeys {
		fmt.Println(m.count)
	}

	counts := make([]int, len(monkeys))
	for i, m := range monkeys {
		counts[i] = m.count
	}
	sort.Ints(counts)
	fmt.Println(counts[len(counts)-1] * counts[len(counts)-2])
}

func parseMonkeys(lines []string) []*monkey {
	var monkeys []*monkey
	var m *monkey
	for _, l := range lines {
		l := strings.TrimSpace(l)
		switch {
		case l == "":
			monkeys = append(monkeys, m)
			m = nil
		case strings.HasPrefix(l, "Monkey"):
			m = new(monkey)
		case strings.HasPrefix(l, "Starting items:"):
			_, nlist, _ := strings.Cut(l, ":")
			for _, s := range strings.Split(nlist, ", ") {
				n := mustParseInt(strings.TrimSpace(s))
				m.items = append(m.items, n)
			}
		case strings.HasPrefix(l, "Operation:"):
			_, eq, _ := strings.Cut(l, "=")
			parts := strings.Split(strings.TrimSpace(eq), " ")
			v0 := m.getOpVal(parts[0])
			v1 := m.getOpVal(parts[2])
			switch parts[1] {
			case "*":
				m.update = makeUpdateFunc(v0, v1, &m.testDivisible, mul)
			case "+":
				m.update = makeUpdateFunc(v0, v1, &m.testDivisible, add)
			default:
				panic(fmt.Sprintf("unknown op %q", parts[1]))
			}
		case strings.HasPrefix(l, "Test:"):
			parts := strings.Split(l, " ")
			m.testDivisible = mustParseInt(parts[len(parts)-1])
		case strings.HasPrefix(l, "If true:"):
			parts := strings.Split(l, " ")
			m.testTrue = mustParseInt(parts[len(parts)-1])
		case strings.HasPrefix(l, "If false:"):
			parts := strings.Split(l, " ")
			m.testFalse = mustParseInt(parts[len(parts)-1])
		}
	}
	monkeys = append(monkeys, m)
	return monkeys
}

func makeUpdateFunc(v0, v1, testDivisible *int, fn func(a, b int, c *int) int) func(item int) int {
	return func(old int) int {
		v0 := v0
		v1 := v1
		if v0 == nil {
			v0 = &old
		}
		if v1 == nil {
			v1 = &old
		}
		v := fn(*v0, *v1, testDivisible)
		return v
	}
}

func (m *monkey) inspectItems(monkeys []*monkey) {
	for len(m.items) > 0 {
		m.count++
		item := m.update(m.items[0]) // % m.testDivisible
		if item < m.items[0] {
			panic(fmt.Sprintf("overflow: %d, %d", item, m.items[0]))
		}
		m.items = m.items[1:]
		var next *monkey
		if item%m.testDivisible == 0 {
			next = monkeys[m.testTrue]
		} else {
			next = monkeys[m.testFalse]
		}
		next.items = append(next.items, item)
	}
}

func (m *monkey) getOpVal(s string) *int {
	if s == "old" {
		return nil
	}
	n := mustParseInt(s)
	return &n
}

func mul(a, b int, testDivisible *int) int {
	return (a * b) // % (*testDivisible)
}

func add(a, b int, testDivisible *int) int {
	return (a + b) // % (*testDivisible)
}
