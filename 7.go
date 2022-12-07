package main

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type fsNode struct {
	path     string
	selfSize int
	cumSize  int
	parent   *fsNode
	children []*fsNode
}

func (n *fsNode) String() string {
	var sb strings.Builder
	sb.WriteString(filepath.Base(n.path))
	sb.WriteByte(' ')
	if n.selfSize == 0 {
		sb.WriteString(fmt.Sprintf("(dir, size=%d)", n.cumSize))
	} else {
		sb.WriteString(fmt.Sprintf("(file, size=%d)", n.selfSize))
	}
	return sb.String()
}

type fsWalker struct {
	cwd  *fsNode
	root *fsNode
}

func day7() {
	lines := mustReadFileLines("7.txt")
	w := newWalker()
	w.processInput(lines)
	w.root.updateSize()

	unused := 70_000_000 - w.root.cumSize
	part1 := 0
	part2 := 0
	var node *fsNode
	nodes := []*fsNode{w.root}
	for len(nodes) > 0 {
		node, nodes = nodes[0], nodes[1:]
		if node.selfSize == 0 && node.cumSize <= 100000 {
			part1 += node.cumSize
		}
		if node.selfSize == 0 && node.cumSize+unused >= 30_000_000 {
			if part2 == 0 || node.cumSize < part2 {
				part2 = node.cumSize
			}
		}
		nodes = append(nodes, node.children...)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func newWalker() *fsWalker {
	w := &fsWalker{root: &fsNode{path: "/"}}
	w.cwd = w.root
	return w
}

func (w *fsWalker) processInput(lines []string) {
	for i := 0; i < len(lines); {
		parts := strings.Split(lines[i], " ")
		if parts[0] != "$" {
			log.Fatalf("unexpected non-command line: %q", lines[i])
		}
		switch parts[1] {
		case "cd":
			switch parts[2] {
			case "/":
				w.cwd = w.root
			case "..":
				if w.cwd == w.root {
					panic("cannot move out of /")
				}
				w.cwd = w.cwd.parent
			default:
				w.cwd = w.addChild(parts[2])
			}
			i++
		case "ls":
			for i = i + 1; i < len(lines) && lines[i][0] != '$'; i++ {
				parts := strings.Split(lines[i], " ")
				if parts[0] == "dir" {
					w.addChild(parts[1])
					continue
				}
				size, err := strconv.Atoi(parts[0])
				if err != nil {
					log.Fatalf("error parsing size from line: %q", lines[i])
				}
				child := w.addChild(parts[1])
				child.selfSize = size
			}
		default:
			log.Fatalf("unknown command %q", parts[1])
		}
	}
}

func (w *fsWalker) addChild(name string) *fsNode {
	path := filepath.Join(w.cwd.path, name)
	for _, node := range w.cwd.children {
		if node.path == path {
			return node
		}
	}
	node := &fsNode{
		path:   path,
		parent: w.cwd,
	}
	w.cwd.children = append(w.cwd.children, node)
	sort.Slice(w.cwd.children, func(i, j int) bool {
		return w.cwd.children[i].path < w.cwd.children[j].path
	})
	return node
}

func (n *fsNode) updateSize() {
	if n.selfSize > 0 {
		n.cumSize = n.selfSize
		return
	}
	for _, child := range n.children {
		child.updateSize()
		n.cumSize += child.cumSize
	}
}

func (w *fsWalker) print() {
	w.printNode(w.root, 0)
}

func (w *fsWalker) printNode(node *fsNode, depth int) {
	for i := 0; i < depth*2; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("- %s\n", node)
	for _, child := range node.children {
		w.printNode(child, depth+1)
	}
}
