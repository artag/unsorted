package main

import (
	"fmt"
	"os"
	"sevseg"
)

func main() {
	str, err := sevseg.GetSegmentsInRow([]string{"20", ".", "05", ".", "2024"}, 2)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Example 1:")
	fmt.Println(str)

	rows, err := sevseg.GetSegmentsInRows([]string{"23", ":", "59", ":", "58"}, 2)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Example 2:")
	fmt.Print(rows[0])
	fmt.Print(rows[1])
	fmt.Print(rows[2])
}
