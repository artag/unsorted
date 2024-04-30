package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tiagomelo/go-clipboard/clipboard"
)

type Command int

const (
	Encrypt Command = iota
	Decrypt
)

var (
	in = bufio.NewReader(os.Stdin)
)

// var chars = []rune{
// 	'A', 'B', 'C', 'D', 'E', 'F', '0', 'G',
// 	'H', 'I', 'J', 'K', 'L', '1', 'M', 'N',
// 	'O', 'P', 'Q', 'R', '2', 'S', 'T', 'U',
// 	'V', 'W', 'X', '3', 'Y', 'Z', 'А', 'Б',
// 	'В', 'Г', '4', 'Д', 'Е', 'Ё', 'Ж', 'З',
// 	'И', '5', 'Й', 'К', 'Л', 'М', 'Н', 'О',
// 	'6', 'П', 'Р', 'С', 'Т', 'У', 'Ф', '7',
// 	'Х', 'Ц', 'Ч', 'Ш', 'Щ', '8', 'Ъ', 'Ы',
// 	'Ь', 'Э', 'Ю', '9', 'Я'}

var chars = "ABCDEF0GHIJKL1MNOPQR2STUVWX3YZАБВГ4ДЕЁЖЗИ5ЙКЛМНО6ПРСТУФ7ХЦЧШЩ8ЪЫЬЭЮ9Я"

func main() {
	sym := NewSymbols(chars)

	fmt.Println("Caesar Cipher")
	cmd := enterCommand(in)
	key := enterKey(in, sym)
	txt := enterText(in, cmd)
	output := executeCommand(txt, key, cmd, sym)
	fmt.Println(output)
	copyToClipboard(output, cmd)
}

func enterCommand(in *bufio.Reader) Command {
	for {
		fmt.Println("Do you want to (e)ncrypt or (d)ecrypt?")
		fmt.Print("> ")

		input, err := in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if err != nil {
			continue
		}

		inputStr := strings.TrimSuffix(input, "\n")
		str := strings.ToLower(inputStr)
		if str == "e" {
			return Encrypt
		}
		if str == "d" {
			return Decrypt
		}
	}
}

func enterKey(in *bufio.Reader, symbols Symbols) int {
	for {
		fmt.Printf("Please enter the key (0 to %d) to use.\n", symbols.Length)
		fmt.Print("> ")

		input, err := in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if err != nil {
			continue
		}

		inputStr := strings.TrimSuffix(input, "\n")
		num, err := strconv.Atoi(inputStr)
		if err != nil {
			continue
		}
		if num < 0 || num > symbols.Length {
			continue
		}

		return num
	}
}

func enterText(in *bufio.Reader, cmd Command) string {
	for {
		if cmd == Encrypt {
			fmt.Println("Enter the message to encrypt.")
		}
		if cmd == Decrypt {
			fmt.Println("Enter the message to decrypt.")
		}
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

func executeCommand(txt string, key int, cmd Command, symbols Symbols) string {
	input := strings.Split(strings.ToUpper(txt), "")
	var output strings.Builder

	for i := 0; i < len(input); i++ {
		ch := input[i]
		idx, foundIdx := symbols.GetIndex(ch)
		if foundIdx {
			newIdx := getNewSymbolIndex(idx, key, symbols, cmd)
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

func getNewSymbolIndex(index int, key int, symbols Symbols, cmd Command) int {
	if cmd == Encrypt {
		newIndex := index + key
		if newIndex < symbols.Length {
			return newIndex
		}
		return index + key - symbols.Length
	}

	if cmd == Decrypt {
		newIndex := index - key
		if newIndex >= 0 {
			return newIndex
		}
		return symbols.Length + index - key
	}

	return index
}

func copyToClipboard(output string, cmd Command) {
	c := clipboard.New()
	if err := c.CopyText(output); err != nil {
		return
	}

	if cmd == Encrypt {
		fmt.Println("Full encrypted text copied to clipboard.")
	}
	if cmd == Decrypt {
		fmt.Println("Full decrypted text copied to clipboard.")
	}
}
