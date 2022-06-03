package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Please input your desired pyramid rows\n")
		os.Exit(1)
	}

	var (
		rows int
		err  error
	)
	rows, err = strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("Need proper number argument\n")
		fmt.Println(err.Error())
		os.Exit(2)
	}

	fmt.Printf("We will be printing pyramid with %d rows\n", rows)
	pyramid(rows)
}
