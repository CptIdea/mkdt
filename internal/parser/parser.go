package parser

import (
	"fmt"
	"strings"
)

type Node struct {
	Name     string
	IsDir    bool
	Depth    int
	Children []*Node
}

func Parse(lines []string) (*Node, error) {
	root := &Node{Depth: -1, IsDir: true}
	stack := []*Node{root}

	for i, line := range lines {
		depth := calculateDepth(line)
		name := strings.TrimSpace(line)
		isDir := strings.HasSuffix(name, "/")

		// remove trailing slash
		if isDir {
			name = strings.TrimSuffix(name, "/")
		}

		// find parent
		for len(stack) > 0 && stack[len(stack)-1].Depth >= depth {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			return nil, fmt.Errorf("line %d: invalid indentation", i+1)
		}

		node := &Node{
			Name:  name,
			IsDir: isDir,
			Depth: depth,
		}

		stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, node)
		stack = append(stack, node)
	}
	return root, nil
}

func calculateDepth(line string) int {
	tabs := 0
	for _, c := range line {
		if c == '\t' {
			tabs++
		} else {
			break
		}
	}
	return tabs
}
