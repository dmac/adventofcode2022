package main

import (
	"fmt"
)

func day3() {
	lines := mustReadFileLines("3.txt")
	sum := 0
	for _, line := range lines {
		line := []byte(line)
		common := commonItems(line[:len(line)/2], line[len(line)/2:])
		sum += priority(common[0])
	}
	fmt.Println(sum)

	sum = 0
	for i := 0; i < len(lines)-2; i += 3 {
		common := commonItems([]byte(lines[i]), []byte(lines[i+1]))
		common = commonItems(common, []byte(lines[i+2]))
		sum += priority(common[0])
	}
	fmt.Println(sum)
}

func commonItems(s0, s1 []byte) []byte {
	s := make(map[byte]struct{})
	var common []byte
	for _, b := range s0 {
		s[b] = struct{}{}
	}
	for _, b := range s1 {
		if _, ok := s[b]; ok {
			common = append(common, b)
		}
	}
	return common
}

func priority(b byte) int {
	if 'a' <= b && b <= 'z' {
		return int(b - 'a' + 1)
	}
	if 'A' <= b && b <= 'Z' {
		return int(b - 'A' + 27)
	}
	panic("bad character")
}
