package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func day13() {
	lines := mustReadFileLines("13.txt")
	var packets []interface{}
	for _, l := range lines {
		if l == "" {
			continue
		}
		var packet interface{}
		if err := json.Unmarshal([]byte(l), &packet); err != nil {
			panic(err)
		}
		packets = append(packets, packet)
	}
	part1 := 0
	for i := 0; i < len(packets)-1; i += 2 {
		p0 := packets[i]
		p1 := packets[i+1]
		if ordered, _ := compareValues(p0, p1); ordered {
			part1 += i/2 + 1
		}
	}
	fmt.Println(part1)

	for _, s := range []string{"[[2]]", "[[6]]"} {
		var div interface{}
		if err := json.Unmarshal([]byte(s), &div); err != nil {
			panic(err)
		}
		packets = append(packets, div)
	}
	sort.Slice(packets, func(i, j int) bool {
		ordered, _ := compareValues(packets[i], packets[j])
		return ordered
	})
	part2 := 1
	for i, p := range packets {
		s := fmt.Sprintf("%v", p)
		if s == "[[2]]" || s == "[[6]]" {
			part2 *= i + 1
		}
	}
	fmt.Println(part2)
}

func compareValues(v0, v1 interface{}) (ordered bool, ok bool) {
	n0, n0ok := v0.(float64)
	n1, n1ok := v1.(float64)
	if n0ok && n1ok {
		if n0 < n1 {
			return true, true
		}
		if n0 > n1 {
			return false, true
		}
	}
	if !n0ok && !n1ok {
		l0 := v0.([]interface{})
		l1 := v1.([]interface{})
		for i := range l0 {
			if i >= len(l1) {
				return false, true
			}
			ordered, ok := compareValues(l0[i], l1[i])
			if ok {
				return ordered, true
			}
		}
		if len(l0) < len(l1) {
			return true, true
		}
	}
	if n0ok && !n1ok {
		return compareValues([]interface{}{n0}, v1)
	}
	if !n0ok && n1ok {
		return compareValues(v0, []interface{}{n1})
	}
	return false, false
}
