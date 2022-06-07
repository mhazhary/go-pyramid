package main

import (
	"fmt"
	"strings"
)

func pyramid(rows int) {
	var stack = make([]string, 2*rows-1)
	for i := range stack {
		stack[i] = " "
	}
	var k int
	for i := rows - 1; i >= 0; i-- {
		stack[i] = "*"
		stack[i+k] = "*"
		fmt.Println(strings.Join(stack, " "))
		k += 2
	}
}
