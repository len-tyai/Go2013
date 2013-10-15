package main

import (
	"strings"
)

var stack []string

func parsePath(path string) string {
	stack = []string{}
	pathParts := strings.Split(path, "/")
	for _, part := range pathParts {
		if part == ".." {
			pop()
		} else if part != "." && part != "" {
			push(part)
		}
	}

	result := "/"
	for i := 0; i < len(stack); i++ {
		result += stack[i] + "/"
	}

	return result
}

func push(path string) {
	stack = append(stack, path)
}

func pop() {
	if len(stack) > 0 {
		stack = stack[0 : len(stack)-1]
	}
}
