package terminal

import (
	"github.com/rivo/tview"
)

type menuOption string

// Definiere Konstanten für die verschiedenen Optionen
const (
	playOption  menuOption = "PLAY"
	exitOption  menuOption = "EXIT"
	rulesOption menuOption = "RULES"
)

var menuOptions = struct {
	selectedItem menuOption
}{}

func RenderMenu() {
	menuPage, menu := getMenuPage()

	App.SetRoot(menuPage, true).SetFocus(menu)
}

func exit() {
	menuOptions.selectedItem = exitOption

	doneFunc := func(buttonIndex int, buttonLabel string) {
		switch buttonLabel {
		case "No":
			RenderMenu()
			break
		case "Yes":
			App.Stop()
			break
		}
	}

	buttonLabels := []string{"Yes", "No"}

	title := "Quit the game"

	text := "Do you really want to quit Kniffel?"

	showModal(text, title, buttonLabels, doneFunc)
}

func rules() {
	menuOptions.selectedItem = rulesOption

	doneFunc := func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Cancel" {
			RenderMenu()
		}
	}

	buttonLabels := []string{"Cancel"}

	title := "Rules"

	text := `1. Each player takes turns rolling five dice.
2. On a player's turn, they may roll the dice up to three times.
   - After each roll, the player can choose to keep any dice and re-roll the rest.
3. The player must choose one scoring category for their turn after their final roll.
4. Each category can only be used once per game.
5. Scoring categories:
   - [::b]Aces (Ones)[::B]: Sum of all dice showing the number 1.
   - [::b]Twos[::B]: Sum of all dice showing the number 2.
   - [::b]Threes[::B]: Sum of all dice showing the number 3.
   - [::b]Fours[::B]: Sum of all dice showing the number 4.
   - [::b]Fives[::B]: Sum of all dice showing the number 5.
   - [::b]Sixes[::B]: Sum of all dice showing the number 6.
   - [::b]Three of a Kind[::B]: Sum of all dice if at least three dice show the same number.
   - [::b]Four of a Kind[::B]: Sum of all dice if at least four dice show the same number.
   - [::b]Full House[::B]: 25 points for a three of a kind and a pair.
   - [::b]Small Straight[::B]: 30 points for a sequence of four consecutive numbers.
   - [::b]Large Straight[::B]: 40 points for a sequence of five consecutive numbers.
   - [::b]Yahtzee[::B]: 50 points for five dice showing the same number.
   - [::b]Chance[::B]: Sum of all dice, no specific requirements.
6. The game ends after all players have filled all categories.
7. The player with the highest total score wins.`

	showModal(text, title, buttonLabels, doneFunc)
}

func getMenuPage() (*tview.Flex, *tview.List) {

	logo := `
┌───────────────────────────────────────────────────────────┐
│ [#0068ac::l] __    __            __   ______    ______            __ [white::L] │
│ [#0068ac::l]|  \  /  \          |  \ /      \  /      \          |  \[white::L] │
│ [#0068ac::l]| $$ /  $$ _______   \$$|  $$$$$$\|  $$$$$$\ ______  | $$[white::L] │
│ [#0068ac::l]| $$/  $$ |       \ |  \| $$_  \$$| $$_  \$$/      \ | $$[white::L] │
│ [#0068ac::l]| $$  $$  | $$$$$$$\| $$| $$ \    | $$ \   |  $$$$$$\| $$[white::L] │
│ [#0068ac::l]| $$$$$\  | $$  | $$| $$| $$$$    | $$$$   | $$    $$| $$[white::L] │
│ [#0068ac::l]| $$ \$$\ | $$  | $$| $$| $$      | $$     | $$$$$$$$| $$[white::L] │
│ [#0068ac::l]| $$  \$$\| $$  | $$| $$| $$      | $$      \$$     \| $$[white::L] │
│ [#0068ac::l] \$$   \$$ \$$   \$$ \$$ \$$       \$$       \$$$$$$$ \$$[white::L] │
└───────────────────────────────────────────────────────────┘
`

	textview := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText(logo)

	menu := tview.NewList().
		AddItem("Play", "Start a new game", 'p', renderUsersForm).
		AddItem("Rules", "View the game rules", 'r', rules).
		AddItem("Exit", "Exit the application", 'e', exit)

	switch menuOptions.selectedItem {
	case playOption:
		menu.SetCurrentItem(0)
	case rulesOption:
		menu.SetCurrentItem(1)
	case exitOption:
		menu.SetCurrentItem(2)
	}

	verticalCenter := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(tview.NewBox(), 0, 4, false).
		AddItem(menu, 0, 1, true).
		AddItem(nil, 0, 4, false)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 15, 1, false).
		AddItem(textview, 15, 2, false).
		AddItem(verticalCenter, 0, 3, true)

	return flex, menu
}

func showModal(text string, title string, buttonLabels []string, doneFunc func(buttonIndex int, buttonLabel string)) {
	m := tview.NewModal().
		SetText(text).
		AddButtons(buttonLabels)

	m.SetDoneFunc(doneFunc)

	m.SetTitle(title).SetTitleAlign(tview.AlignCenter)

	App.SetRoot(m, true).SetFocus(m)
}
