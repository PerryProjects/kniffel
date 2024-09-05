package diceCupModel

import (
	"cli-kniffel/pkg/cli"
	"cli-kniffel/pkg/model/dice"
	"fmt"
	"math/rand"
	"time"
)

const numDice = 5

type DiceCup struct {
	dices [numDice]*dice.Dice
}

func (diceCup DiceCup) GetDices() [numDice]*dice.Dice {
	return diceCup.dices
}

func NewDiceCup() *DiceCup {
	diceCup := DiceCup{}
	cli.HideCursor()
	defer cli.ShowCursor()
	for i := 0; i < rand.Intn(25-10)+10; i++ {
		cli.ClearScreen()
		for j := 0; j < numDice; j++ {
			newDice := dice.NewDice()
			newDice.Roll()
			diceCup.dices[j] = newDice
		}
		PrintDices(diceCup.dices)
		//fmt.Printf("Dice %d: %d\n", results[j].Number, results[j].Value)
		time.Sleep(100 * time.Millisecond)
	}
	return &diceCup
}

func PrintDices(dices [numDice]*dice.Dice) {
	for line := 0; line < 10; line++ {
		for i := 0; i < 5; i++ {
			if line != 9 {
				fmt.Print(dice.Drawings[dices[i].CurrentValue][line] + "    ")
			} else {
				if i == 0 {
					fmt.Printf("     Dice %d", i+1)
				} else {
					fmt.Printf("              Dice %d", i+1)
				}
			}
		}
		fmt.Println()
	}
	fmt.Println("")
}
