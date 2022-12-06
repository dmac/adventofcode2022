package main

import "fmt"

func day6() {
	lines := mustReadFileLines("6.txt")
	input := lines[0]
	fmt.Println(detectMarker(input, 4))
	fmt.Println(detectMarker(input, 14))
}

func detectMarker(input string, size int) int {
	buf := make([]byte, size)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(buf)-1; j++ {
			buf[j] = buf[j+1]
		}
		buf[len(buf)-1] = input[i]
		if i < len(buf) {
			continue
		}
		s := make(map[byte]struct{})
		for j := 0; j < len(buf); j++ {
			s[buf[j]] = struct{}{}
		}
		if len(s) == len(buf) {
			return i + 1
		}
	}
	return -1
}
