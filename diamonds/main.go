package main

import (
	"fmt"
	"strings"
)

type Size int

func main() {
	sizes := []int{1, 2, 3, 4, 5, 6}
	for _, size := range sizes {
		displayOutlineDiamond(size)
		fmt.Println()
		displayFilledDiamond(size)
		fmt.Println()
	}
}

func displayOutlineDiamond(size int) {
	var sb strings.Builder

	// Upper half
	for i := 0; i < size; i++ {
		for j := 0; j <= size*2; j++ {
			if j == size-1-i {
				sb.WriteRune('/')
			} else if j == size+i {
				sb.WriteRune('\\')
			} else if j == size*2-1 {
				sb.WriteRune(' ')
			} else if j == size*2 {
				sb.WriteRune('\n')
			} else {
				sb.WriteRune(' ')
			}
		}
	}

	// Bottom half
	for i := 0; i < size; i++ {
		for j := 0; j <= size*2; j++ {
			if j == i {
				sb.WriteRune('\\')
			} else if j == size*2-1-i {
				sb.WriteRune('/')
			} else if j == size*2 {
				sb.WriteRune('\n')
			} else {
				sb.WriteRune(' ')
			}
		}
	}

	fmt.Print(sb.String())
}

func displayFilledDiamond(size int) {
	var sb strings.Builder

	// Upper half
	for i := 0; i < size; i++ {
		for j := 0; j <= size*2; j++ {
			if size-i-1 <= j && j < size {
				sb.WriteRune('/')
			} else if size <= j && j <= size+i {
				sb.WriteRune('\\')
			} else if j == size*2-1 {
				sb.WriteRune(' ')
			} else if j == size*2 {
				sb.WriteRune('\n')
			} else {
				sb.WriteRune(' ')
			}
		}
	}

	// Bottom half
	for i := 0; i < size; i++ {
		for j := 0; j <= size*2; j++ {
			if i <= j && j < size {
				sb.WriteRune('\\')
			} else if size <= j && j <= size*2-i-1 {
				sb.WriteRune('/')
			} else if j == size*2 {
				sb.WriteRune('\n')
			} else {
				sb.WriteRune(' ')
			}
		}
	}

	fmt.Print(sb.String())
}
