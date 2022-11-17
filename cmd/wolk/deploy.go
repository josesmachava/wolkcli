package wolk

import (
	"github.com/spf13/cobra"
	"wolk/pkg/wolk"
)

var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"deploy"},
	Short:   "Deploy a stactic website to wolk",

	Run: func(cmd *cobra.Command, args []string) {
		wolk.ReadCurrentDir()
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
