package terminal

import (
	"cli-kniffel/internal/model/players"
	"github.com/rivo/tview"
	"strconv"
)

func renderUsersForm() {
	// Create main grid for players
	main := tview.NewGrid().
		SetRows(10, 0, 0, 0, 0, 11).
		SetColumns(30, 0, 30)

	// Create input fields for each player
	for i, player := range players.Players {
		inputField := tview.NewInputField().
			SetText(player.Name).
			SetChangedFunc(func(text string) {
				// Update the player's name whenever the input changes
				player.Name = text
			}).
			SetLabel("Player " + strconv.Itoa(i+1) + ": ")

		main.
			AddItem(inputField, i+1, 1, 1, 1, 0, 0, i == 0)
	}

	// Create grid layout
	grid := tview.NewGrid().
		SetRows(0, 30, 0).
		SetColumns(0, 100, 0).
		AddItem(main, 1, 1, 1, 1, 0, 100, true)

	// Set the root of the application
	App.SetRoot(grid, true)

}
