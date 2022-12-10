package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) > 2 {
		log.Fatal("Usage: adventofcode2022 [day]")
	}
	if len(os.Args) == 2 {
		day, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		switch day {
		case 1:
			day1()
		case 2:
			day2()
		case 3:
			day3()
		case 4:
			day4()
		case 5:
			day5()
		case 6:
			day6()
		case 7:
			day7()
		case 8:
			day8()
		case 9:
			day9()
		case 10:
			day10()
		case 11:
		case 12:
		case 13:
		case 14:
		case 15:
		case 16:
		case 17:
		case 18:
		case 19:
		case 20:
		case 21:
		case 22:
		case 23:
		case 24:
		case 25:
		}
	}
}

func mustReadFileLines(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
