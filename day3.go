package main

import (
	"fmt"
	"log"
)

func day3() {
	scanner := makeFileScanner("3.txt")
	sum := 0
	var lines [][]byte
	for scanner.Scan() {
		line := scanner.Bytes()
		common := commonItems(line[:len(line)/2], line[len(line)/2:])
		sum += priority(common[0])
		lines = append(lines, append([]byte{}, line...))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)

	sum = 0
	for i := 0; i < len(lines)-2; i += 3 {
		common := commonItems(lines[i], lines[i+1])
		common = commonItems(common, lines[i+2])
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
