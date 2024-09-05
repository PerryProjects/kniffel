package cli

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScreen() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func HideCursor() {
	fmt.Print("\033[?25l") // Cursor verstecken
}

func ShowCursor() {
	fmt.Print("\033[?25h") // Cursor wieder anzeigen
}
