package main

import (
	"fmt"
	"strings"
)

type Symbols struct {
	dic1   map[rune]int
	dic2   map[int]rune
	Length int
}

func NewSymbols(chars string) Symbols {
	upper := strings.ToUpper(chars)
	len := len(upper)
	dic1 := make(map[rune]int)
	dic2 := make(map[int]rune)
	for i := 0; i < len; i++ {
		ch := rune(upper[i])
		dic1[ch] = i
		dic2[i] = ch
	}
	fmt.Println()

	return Symbols{
		dic1:   dic1,
		dic2:   dic2,
		Length: len,
	}
}

func (s *Symbols) GetSymbol(index int) (rune, bool) {
	symbol, ok := s.dic2[index]
	return symbol, ok
}

func (s *Symbols) GetIndex(symbol rune) (int, bool) {
	val, ok := s.dic1[symbol]
	return val, ok
}
