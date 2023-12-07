package main

import (
	"fmt"
	"sort"
	"strings"
)

type valve struct {
	name    string
	flow    int
	open    bool
	tunnels []string
}

type valvePathfinder struct {
	maxPressure int
	maxPath     string
	valves      map[string]*valve
}

func day16() {
	valves := make(map[string]*valve)
	lines := mustReadFileLines("16.txt")
	for _, l := range lines {
		parts := strings.Split(l, " ")
		_, ns, _ := strings.Cut(parts[4], "=")
		i := strings.LastIndex(l, "valve")
		i += strings.Index(l[i:], " ")
		v := &valve{
			name:    parts[1],
			flow:    mustParseInt(ns[:len(ns)-1]),
			tunnels: strings.Split(l[i+1:], ", "),
		}
		valves[parts[1]] = v
	}

	vp := &valvePathfinder{valves: valves}
	vp.findMaxPressure([]string{"AA"})
	fmt.Println(vp.maxPressure)

	// 	permutations := makePermutations(valves, nil, "AA")
	// 	fmt.Println(len(permutations))

	// 	os.Exit(0)

	// 	minutes := 0
	// 	flowing := 0
	// 	pressure := 0
	// 	curr := "AA"
	// outer:
	// 	for {
	// 		next := nextBestScore(valves, curr, minutes)
	// 		if next == "" {
	// 			break
	// 		}
	// 		path := findPath(valves, curr, next)
	// 		fmt.Println(path)
	// 		if len(path) == 0 {
	// 			panic("0 length path")
	// 		}
	// 		for _, name := range path {
	// 			curr = name
	// 			minutes++
	// 			pressure += flowing
	// 			// fmt.Printf("minute %d: pressure=%d\n", minutes, pressure)
	// 			if minutes == 30 {
	// 				break outer
	// 			}
	// 		}
	// 		v := valves[next]
	// 		minutes++
	// 		pressure += flowing
	// 		// fmt.Printf("minute %d: pressure=%d\n", minutes, pressure)
	// 		v.open = true
	// 		flowing += v.flow
	// 		if minutes == 30 {
	// 			break outer
	// 		}
	// 	}
	// 	for minutes < 30 {
	// 		minutes++
	// 		pressure += flowing
	// 		// fmt.Printf("minute %d: pressure=%d\n", minutes, pressure)
	// 	}
	// 	fmt.Println(pressure)
	// 	var z uint64
	// 	for i := 0; i < 1e12; i++ {
	// 		z++
	// 	}
	// 	fmt.Println(z)
}

func (vp *valvePathfinder) findMaxPressure(path []string) {
	var rem []string
outer:
	for name, v := range vp.valves {
		if v.flow == 0 {
			continue
		}
		for _, p := range path {
			if name == p {
				continue outer
			}
		}
		rem = append(rem, name)
	}
	sort.Strings(rem)
	for _, name := range rem {
		path := append(path, name)
		vp.findMaxPressure(path)
	}
	// if len(rem) == 0 {
	// 	fmt.Println(path)
	// }
}

func findPath(valves map[string]*valve, start, goal string) []string {
	type path struct {
		v       *valve
		tunnels []string
	}
	paths := []*path{{v: valves[start]}}
	visited := make(map[string]struct{})
	var p *path
	for len(paths) > 0 {
		p, paths = paths[0], paths[1:]
		if p.v.name == goal {
			return p.tunnels
		}
		for _, tunnel := range p.v.tunnels {
			if _, ok := visited[tunnel]; !ok {
				visited[tunnel] = struct{}{}
				tunnels := make([]string, len(p.tunnels)+1)
				for i, t := range p.tunnels {
					tunnels[i] = t
				}
				tunnels[len(tunnels)-1] = tunnel
				next := &path{
					v:       valves[tunnel],
					tunnels: tunnels,
				}
				paths = append(paths, next)
			}
		}
	}
	return nil
}

func highestClosedValve(valves map[string]*valve) string {
	var vs []*valve
	for _, v := range valves {
		vs = append(vs, v)
	}
	sort.Slice(vs, func(i, j int) bool {
		return vs[i].flow >= vs[j].flow
	})
	for _, v := range vs {
		if !v.open && v.flow > 0 {
			return v.name
		}
	}
	return ""
}

func nextBestScore(valves map[string]*valve, start string, minutes int) string {
	type option struct {
		name  string
		score int
	}
	var opts []*option
	for _, v := range valves {
		if v.name == start || v.flow == 0 || v.open {
			continue
		}
		path := findPath(valves, start, v.name)
		if minutes+len(path) > 30 {
			continue
		}
		opt := &option{
			name:  v.name,
			score: v.flow * (30 - minutes - len(path) - 1),
		}
		opts = append(opts, opt)
	}
	sort.Slice(opts, func(i, j int) bool {
		if opts[i].score != opts[j].score {
			return opts[i].score >= opts[j].score
		}
		return opts[i].name < opts[j].name
	})
	if len(opts) > 0 {
		return opts[0].name
	}
	return ""
}

func closedValves(valves map[string]*valve, open []string) []string {
	var vs []*valve
outer:
	for _, v := range valves {
		if v.flow == 0 {
			continue
		}
		for _, n := range open {
			if v.name == n {
				continue outer
			}
		}
		vs = append(vs, v)
	}
	sort.Slice(vs, func(i, j int) bool {
		return vs[i].name < vs[j].name
	})
	var names []string
	for _, v := range vs {
		names = append(names, v.name)
	}
	return names
}

type valvePermutation struct {
	open []bool
	path []string
}

func makePermutations(valves map[string]*valve, perm *valvePermutation, start string) []*valvePermutation {
	var perms []*valvePermutation
	if perm == nil {
		perm0 := &valvePermutation{
			open: []bool{false},
			path: []string{"AA"},
		}
		perm1 := &valvePermutation{
			open: []bool{true},
			path: []string{"AA"},
		}
		perms = append(perms, makePermutations(valves, perm0, start)...)
		perms = append(perms, makePermutations(valves, perm1, start)...)
		return perms
	}
	for _, next := range valves[start].tunnels {
		// for _, name := range perm.path {
		// 	if name == next {
		// 		continue
		// 	}
		// }
		perm0 := &valvePermutation{
			open: append(copyBools(perm.open), false),
			path: append(copyStrings(perm.path), next),
		}
		perm1 := &valvePermutation{
			open: append(copyBools(perm.open), true),
			path: append(copyStrings(perm.path), next),
		}
		perms = append(perms, makePermutations(valves, perm0, next)...)
		perms = append(perms, makePermutations(valves, perm1, next)...)
	}
	return perms
}

func copyBools(s []bool) []bool {
	ss := make([]bool, len(s))
	for i := range s {
		ss[i] = s[i]
	}
	return ss
}

func copyStrings(s []string) []string {
	ss := make([]string, len(s))
	for i := range s {
		ss[i] = s[i]
	}
	return ss
}
