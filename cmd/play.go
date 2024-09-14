package cmd

import (
	"cli-kniffel/internal/terminal"
	"github.com/spf13/cobra"
)

var numDice = 5

// rollCmd represents the roll command
var play = &cobra.Command{
	Use:   "play",
	Short: "Start a new game",
	Long:  `Start a new game of Kniffel. You can play with up to four players.`,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func init() {
	rootCmd.AddCommand(play)
}

func start() {
	terminal.RenderMenu()

	err := terminal.App.EnableMouse(true).Run()
	if err != nil {
		panic(err)
	}
}
