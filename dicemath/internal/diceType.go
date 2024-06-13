package internal

type diceType struct {
	rows  []string
	value int
}

var allDiceTypes = []diceType{
	d1,
	d1,
	d2a,
	d2b,
	d3a,
	d3b,
	d4,
	d4,
	d5,
	d5,
	d6a,
	d6b,
}

var d1 = diceType{
	rows: []string{
		"+-------+",
		"|       |",
		"|   O   |",
		"|       |",
		"+-------+",
	},
	value: 1,
}

var d2a = diceType{
	rows: []string{
		"+-------+",
		"| O     |",
		"|       |",
		"|     O |",
		"+-------+",
	},
	value: 2,
}

var d2b = diceType{
	rows: []string{
		"+-------+",
		"|     O |",
		"|       |",
		"| O     |",
		"+-------+",
	},
	value: 2,
}

var d3a = diceType{
	rows: []string{
		"+-------+",
		"| O     |",
		"|   O   |",
		"|     O |",
		"+-------+",
	},
	value: 3,
}

var d3b = diceType{
	rows: []string{
		"+-------+",
		"|     O |",
		"|   O   |",
		"| O     |",
		"+-------+",
	},
	value: 3,
}

var d4 = diceType{
	rows: []string{
		"+-------+",
		"| O   O |",
		"|       |",
		"| O   O |",
		"+-------+",
	},
	value: 4,
}

var d5 = diceType{
	rows: []string{
		"+-------+",
		"| O   O |",
		"|   O   |",
		"| O   O |",
		"+-------+",
	},
	value: 5,
}

var d6a = diceType{
	rows: []string{
		"+-------+",
		"| O   O |",
		"| O   O |",
		"| O   O |",
		"+-------+",
	},
	value: 6,
}

var d6b = diceType{
	rows: []string{
		"+-------+",
		"| O O O |",
		"|       |",
		"| O O O |",
		"+-------+",
	},
	value: 6,
}
