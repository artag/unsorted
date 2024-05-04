package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	in = bufio.NewReader(os.Stdin)
)

var chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	sym := NewSymbols(chars)
	fmt.Println("Caesar Cipher Hacker")
	msg := enterMessage(in)
	hackMessage(msg, sym)
}

func enterMessage(in *bufio.Reader) string {
	for {
		fmt.Println("Enter the encrypted Caesar cipher message to hack.")
		fmt.Print("> ")

		input, err := in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if err != nil {
			continue
		}
		if len(input) < 1 {
			continue
		}

		return input
	}
}

func hackMessage(msg string, symbols Symbols) {
	for key := 0; key < symbols.Length; key++ {
		output := hackMessageWithKey(msg, key, symbols)
		fmt.Printf("Key #%2d: %s\n", key, output)
	}
}

func hackMessageWithKey(msg string, key int, symbols Symbols) string {
	input := strings.Split(strings.ToUpper(msg), "")
	var output strings.Builder

	for i := 0; i < len(input); i++ {
		ch := input[i]
		idx, foundIdx := symbols.GetIndex(ch)
		if foundIdx {
			newIdx := getNewSymbolIndex(idx, key, symbols)
			newCh, foundCh := symbols.GetSymbol(newIdx)
			if foundCh {
				output.WriteString(newCh)
			} else {
				output.WriteString(ch)
			}
		} else {
			output.WriteString(ch)
		}
	}

	return output.String()
}

func getNewSymbolIndex(index int, key int, symbols Symbols) int {
	newIndex := index - key
	if newIndex >= 0 {
		return newIndex
	}

	return symbols.Length + index - key
}
