package cmd

import (
	"cli-kniffel/model/diceCup"
	"fmt"
	"github.com/spf13/cobra"
)

var numDice = 5

// rollCmd represents the roll command
var rollCmd = &cobra.Command{
	Use:   "roll",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func init() {
	rootCmd.AddCommand(rollCmd)
}

func start() {
	diceCup := diceCupModel.NewDiceCup()

	//clearScreen()

	//results, _ = selectDices(0, results)

	fmt.Printf("You choose %d dices\n", len(diceCup.GetDices()))
}

/*func selectDices(selectedPos int, allDices []*dice) ([]*dice, error) {
	const doneID = 7
	if len(allDices) > 0 && allDices[0].Number != doneID {
		var items = []*dice{
			{
				Number: 7,
				Value:  0,
			},
		}
		allDices = append(items, allDices...)
	}

	templates := &promptui.SelectTemplates{
		Label: `{{if .IsSelected}}
                    ✔
                {{end}} Dice {{ .Number }}: {{ .Value }}`,
		Active:   "→ {{if .IsSelected}} ✔ {{end}}  Dice {{ .Number  }} : {{ .Value  | cyan }}",
		Inactive: "{{if .IsSelected}} ✔ {{end}} Dice {{ .Number }} : {{ .Value | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Select the dice you want to keep",
		Items:     allDices,
		Templates: templates,
		// Start the cursor at the currently selected index
		CursorPos:    selectedPos,
		Size:         7,
		HideSelected: true,
	}

	selectionIdx, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, nil
	}

	chosenItem := allDices[selectionIdx]

	if chosenItem.Number != 7 {
		// If the user selected something other than "Done",
		// toggle selection on this item and run the function again.
		chosenItem.IsSelected = !chosenItem.IsSelected
		_, err := selectDices(selectionIdx, allDices)
		if err != nil {
			return nil, err
		}
	}

	var selectedItems []*dice
	for _, i := range allDices {
		if i.IsSelected {
			selectedItems = append(selectedItems, i)
		}
	}
	return selectedItems, nil
}
*/
