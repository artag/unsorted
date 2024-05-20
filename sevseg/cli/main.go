package main

import (
	"fmt"
	"os"
	"sevseg"
)

func main() {
	str, err := sevseg.GetSevSegStr("20-05-2024.", 4)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(str)
}
