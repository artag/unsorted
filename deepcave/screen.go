package main

type Screen struct {
	rows   []string
	Width  int
	Height int
}

type Printer func(x, y int, char rune)

func NewScreen(width, height int) *Screen {
	return &Screen{
		rows:   make([]string, height-1),
		Width:  width,
		Height: height,
	}
}

func (s *Screen) Pop() string {
	r := s.rows[0]
	s.rows = s.rows[1:]
	return r
}

func (s *Screen) Push(row string) {
	s.rows = append(s.rows, row)
}

func (s *Screen) Display(printer Printer) {
	for y, row := range s.rows {
		for x, char := range row {
			printer(x, y, char)
		}
	}
}
