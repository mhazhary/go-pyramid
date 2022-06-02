package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Need number argument\n")
		return
	}

	var (
		rows int
		err  error
	)
	rows, err = strconv.Atoi(os.Args[1])
	var k int

	if err != nil {
		fmt.Printf("Need proper number argument\n")
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("We will be printing pyramid with %d rows\n", rows)

	for i := 1; i <= rows; i++ {
		k = 0
		for space := 1; space <= rows-i; space++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print("* ")
			k++
			if k == 2*i-1 {
				break
			}
		}
		fmt.Println("")
	}
}
