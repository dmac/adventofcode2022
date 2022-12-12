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
			switch parts[1] {
			case "*":
				if parts[2] == "old" {
					m.update = func(a int) int { return a * a }
				} else {
					v := mustParseInt(parts[2])
					m.update = func(a int) int { return a * v }
				}
			case "+":
				v := mustParseInt(parts[2])
				m.update = func(a int) int { return a + v }
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

func (m *monkey) inspectItems(monkeys []*monkey) {
	// if m == monkeys[0] {
	// 	fmt.Println(m.items)
	// }
	for len(m.items) > 0 {
		m.count++
		item := m.update(m.items[0]) % m.testDivisible
		// if m == monkeys[0] {
		// 	fmt.Println(m.items[0], item, m.testDivisible, item%m.testDivisible == 0)
		// }
		// if item < m.items[0] {
		// 	panic(fmt.Sprintf("overflow: %d, %d", item, m.items[0]))
		// }
		var next *monkey
		if item%m.testDivisible == 0 {
			next = monkeys[m.testTrue]
		} else {
			next = monkeys[m.testFalse]
		}
		next.items = append(next.items, item)
		m.items = m.items[1:]
	}
}
