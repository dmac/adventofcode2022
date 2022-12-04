package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type assignment struct {
	lo0 int
	hi0 int
	lo1 int
	hi1 int
}

func day4() {
	var assns []*assignment
	scanner := makeFileScanner("4.txt")
	for scanner.Scan() {
		pair0, pair1, _ := strings.Cut(scanner.Text(), ",")
		lo0, hi0, _ := strings.Cut(pair0, "-")
		lo1, hi1, _ := strings.Cut(pair1, "-")
		assn := &assignment{
			lo0: mustParseInt(lo0),
			hi0: mustParseInt(hi0),
			lo1: mustParseInt(lo1),
			hi1: mustParseInt(hi1),
		}
		assns = append(assns, assn)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	containCount := 0
	overlapCount := 0
	for _, a := range assns {
		if contains(a.lo0, a.hi0, a.lo1, a.hi1) {
			containCount++
		}
		if overlaps(a.lo0, a.hi0, a.lo1, a.hi1) {
			overlapCount++
		}
	}
	fmt.Println(containCount)
	fmt.Println(overlapCount)
}

func mustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func contains(lo0, hi0, lo1, hi1 int) bool {
	if lo1 >= lo0 && hi1 <= hi0 {
		return true
	}
	if lo0 >= lo1 && hi0 <= hi1 {
		return true
	}
	return false
}

func overlaps(lo0, hi0, lo1, hi1 int) bool {
	if lo1 <= hi0 && hi1 >= lo0 {
		return true
	}
	if lo0 <= hi1 && hi0 >= lo1 {
		return true
	}
	return false
}
