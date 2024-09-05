package diceModel

import (
	"math/rand"
)

var Drawings = map[int][]string{
	1: []string{
		" -------------- ",
		"|              |",
		"|              |",
		"|              |",
		"|       O      |",
		"|              |",
		"|              |",
		"|              |",
		" -------------- ",
	},
	2: []string{
		" -------------- ",
		"|              |",
		"|    O         |",
		"|              |",
		"|              |",
		"|              |",
		"|          O   |",
		"|              |",
		" -------------- ",
	},
	3: []string{
		" -------------- ",
		"|              |",
		"|    O         |",
		"|              |",
		"|       O      |",
		"|              |",
		"|          O   |",
		"|              |",
		" -------------- ",
	},
	4: []string{
		" -------------- ",
		"|              |",
		"|    O     O   |",
		"|              |",
		"|              |",
		"|              |",
		"|    O     O   |",
		"|              |",
		" -------------- ",
	},
	5: []string{
		" -------------- ",
		"|              |",
		"|    O     O   |",
		"|              |",
		"|       O      |",
		"|              |",
		"|    O     O   |",
		"|              |",
		" -------------- ",
	},
	6: []string{
		" -------------- ",
		"|              |",
		"|    O     O   |",
		"|              |",
		"|    O     O   |",
		"|              |",
		"|    O     O   |",
		"|              |",
		" -------------- ",
	},
}

type Dice struct {
	CurrentValue int
	isSelected   bool
}

func NewDice() *Dice {
	return &Dice{
		isSelected: false,
	}
}

func (d *Dice) Roll() {
	d.CurrentValue = rand.Intn(6) + 1
}

func (d *Dice) Select() {
	d.isSelected = !d.isSelected
}
