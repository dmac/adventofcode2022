package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func day1() {
	tot := 0
	var tots []int
	lines := mustReadFileLines("1.txt")
	for _, s := range lines {
		if s == "" {
			tots = append(tots, tot)
			tot = 0
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		tot += n
	}
	tots = append(tots, tot)
	sort.Slice(tots, func(i, j int) bool {
		return tots[i] >= tots[j]
	})
	fmt.Println(tots[0])
	fmt.Println(tots[0] + tots[1] + tots[2])
}
