package terminal

import (
	"cli-kniffel/internal/model/players"
	"github.com/rivo/tview"
	"strconv"
)

func renderUsersForm() {

	addUser := func() {
		players.Players = append(players.Players, players.Player{})
		renderUsersForm()
	}

	deleteUser := func(i int) {
		players.Players = append(players.Players[:i], players.Players[i+1:]...)
		renderUsersForm()
	}

	flexRows := tview.NewFlex().SetDirection(tview.FlexRow)
	formColumn := tview.NewFlex().
		SetDirection(tview.FlexColumn)

	for i, player := range players.Players {
		formColumn.
			AddItem(tview.NewBox(), 0, 2, false)

		formColumn.
			AddItem(tview.NewInputField().
				SetLabel("Player "+strconv.Itoa(i+1)).
				SetText(player.Name),
				50, 1, false).
			AddItem(tview.NewBox(), 5, 1, false)

		//if i != 0 {
		formColumn.
			AddItem(tview.NewButton("Delete").
				SetSelectedFunc(func() {
					deleteUser(i)
				}),
				10, 1, false)
		//}

		formColumn.
			AddItem(tview.NewBox(), 0, 2, false)

	}
	flexRows.AddItem(formColumn, 1, 0, false)
	flexRows.AddItem(tview.NewBox(), 1, 0, false)

	b := tview.NewButton("Start Game")

	addUserButton := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 2, false).
		AddItem(tview.NewButton("Add user").
			SetSelectedFunc(addUser), 0, 1, true).
		AddItem(nil, 2, 1, false).
		AddItem(b, 0, 1, true).
		AddItem(nil, 0, 2, false)

	// Füge etwas Abstand oberhalb des Buttons hinzu
	flexRows.AddItem(tview.NewBox(), 2, 0, false)

	// Füge den "Add User"-Button zur Liste hinzu
	flexRows.AddItem(addUserButton, 1, 0, false)

	App.SetRoot(flexRows, true)
}
