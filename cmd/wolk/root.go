package wolk

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wolk",
	Short: "The wolkclt command is your primary way to interact with wolk's Global Application Platform",
	Long:  `The wolkclt command is your primary way to interact with wolk's Global Application Platform`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
