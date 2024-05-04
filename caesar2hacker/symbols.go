package main

import (
	"fmt"
	"strings"
)

type Symbols struct {
	dic1   map[string]int
	dic2   map[int]string
	Length int
}

func NewSymbols(chars string) Symbols {
	split := strings.Split(chars, "")
	len := len(split)
	dic1 := make(map[string]int)
	dic2 := make(map[int]string)
	for i := 0; i < len; i++ {
		ch := split[i]
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

func (s *Symbols) GetSymbol(index int) (string, bool) {
	symbol, ok := s.dic2[index]
	return symbol, ok
}

func (s *Symbols) GetIndex(symbol string) (int, bool) {
	val, ok := s.dic1[symbol]
	return val, ok
}
