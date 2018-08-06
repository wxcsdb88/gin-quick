package cmd

import (
	"sync"

	"github.com/spf13/cobra"
)

var (
	apiConfigFile *string
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the api",
	Long: `usage example:
	api(.exe) start -c config/app.toml
	start the api`,
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup

		wg.Add(1)
		wg.Wait()

	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	apiConfigFile = startCmd.Flags().StringP("config", "c", "", "api config file (required)")
	startCmd.MarkFlagRequired("config")
}
