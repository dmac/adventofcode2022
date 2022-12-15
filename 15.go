package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func day15() {
	const (
		// filename  = "15small.txt"
		// targetRow = 10
		// tuningMax = 20
		filename  = "15.txt"
		targetRow = 2000000
		tuningMax = 4000000
	)
	grid := make(map[point]byte)
	sensors := make(map[point]point) // sensor -> beacon
	beacons := make(map[point]point) // beacon -> sensor
	lines := mustReadFileLines(filename)
	for _, l := range lines {
		if strings.HasPrefix(l, "#") {
			continue
		}
		i := strings.Index(l, "x=")
		j := i + strings.Index(l[i:], ",")
		sx := mustParseInt(l[i+2 : j])

		i = strings.Index(l, "y=")
		j = i + strings.Index(l[i:], ":")
		sy := mustParseInt(l[i+2 : j])

		i = strings.LastIndex(l, "x=")
		j = i + strings.Index(l[i:], ",")
		bx := mustParseInt(l[i+2 : j])

		i = strings.LastIndex(l, "y=")
		j = len(l)
		by := mustParseInt(l[i+2 : j])

		sensor := point{sx, sy}
		beacon := point{bx, by}
		grid[sensor] = 'S'
		grid[beacon] = 'B'
		sensors[sensor] = beacon
		beacons[beacon] = sensor
	}
	minX, maxX := math.MaxInt, math.MinInt
	for sensor, beacon := range sensors {
		dist := manhattanDist(sensor, beacon)
		if sensor.x-dist < minX {
			minX = sensor.x - dist
		}
		if sensor.x+dist > maxX {
			maxX = sensor.x + dist
		}
	}
	empty, _ := checkRow(grid, sensors, beacons, targetRow, minX, maxX)
	fmt.Println(empty)

	minX = 0
	maxX = 20
	for y := 0; y < tuningMax; y++ {
		_, segments := checkRow(grid, sensors, beacons, y, 0, tuningMax)
		if len(segments) > 0 {
			fmt.Println(segments[0].min*4000000 + y)
			break
		}
	}
}

type segment struct {
	min int
	max int
}

func checkRow(grid map[point]byte, sensors, beacons map[point]point, row int, minX, maxX int) (empty int, segments []segment) {
	segments = []segment{{minX, maxX}}
	for sensor, beacon := range sensors {
		beaconDist := manhattanDist(sensor, beacon)
		yRange := segment{
			sensor.y - beaconDist,
			sensor.y + beaconDist,
		}
		if yRange.min > row || yRange.max < row {
			continue
		}
		halfCut := beaconDist - abs(sensor.y-row)
		cut := segment{
			sensor.x - halfCut,
			sensor.x + halfCut,
		}
		var newSegments []segment
		for _, seg := range segments {
			if cut.min > seg.max || cut.max < seg.min {
				// no intersection
				newSegments = append(newSegments, seg)
				continue
			}
			if cut.min <= seg.min && cut.max <= seg.max {
				// left edge
				newSeg := segment{cut.max + 1, seg.max}
				if newSeg.min <= newSeg.max {
					newSegments = append(newSegments, newSeg)
				}
				continue
			}
			if cut.min >= seg.min && cut.max >= seg.max {
				// right edge
				newSeg := segment{seg.min, cut.min - 1}
				if newSeg.min <= newSeg.max {
					newSegments = append(newSegments, newSeg)
				}
				continue
			}
			if cut.min > seg.min && cut.max < seg.max {
				// middle
				newSegments = append(newSegments, segment{seg.min, cut.min - 1})
				newSegments = append(newSegments, segment{cut.max + 1, seg.max})
				continue
			}
			if cut.min <= seg.min && cut.max >= seg.max {
				// outside
				continue
			}
			panic("unhandled cut")
		}
		segments = newSegments
		sort.Slice(segments, func(i, j int) bool {
			return segments[i].min < segments[j].min
		})
	}
	empty = maxX - minX + 1
	for _, seg := range segments {
		empty -= (seg.max - seg.min + 1)
	}
	for beacon := range beacons {
		if beacon.y != row {
			continue
		}
		for i := 0; i < len(segments)-1; i++ {
			seg := segment{segments[i].max + 1, segments[i+1].min - 1}
			if beacon.x >= seg.min && beacon.x <= seg.max {
				empty--
			}
		}
	}
	return empty, segments
}

func manhattanDist(p0, p1 point) int {
	return abs(p0.x-p1.x) + abs(p0.y-p1.y)
}
