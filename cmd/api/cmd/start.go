package cmd

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"github.com/wxcsdb88/gin-quick/api"

	"github.com/wxcsdb88/gin-quick/log/logruslogger"

	"github.com/spf13/cobra"

	"github.com/wxcsdb88/gin-quick/cmd"
	"github.com/wxcsdb88/gin-quick/config"
)

var (
	apiConfigFile *string
	versionFlag   *bool
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the api",
	Long: `usage example:
	api(.exe) start -c config/app.toml
	start the api`,
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		globalConfig := config.LoadConfig(*apiConfigFile)
		relaFilePath, _ := filepath.Abs(*apiConfigFile)
		fmt.Printf("load config file: %v\n", relaFilePath)
		fmt.Printf("config: %#v\n\n", globalConfig)

		logger := logruslogger.GetLoggerWithCaller("api", true, true, globalConfig)
		logger.Info("start api at %v", time.Now().Local())

		apiServer, err := api.New(globalConfig)
		if err != nil {
			logger.Error(err.Error())
			return
		}
		if err := apiServer.Start(); err != nil {
			logger.Error("api server run failed %v", err)
			return
		}

		wg.Add(1)
		wg.Wait()

	},
}

func init() {
	// add version cmd
	rootCmd.AddCommand(cmd.VersionCmd)
	versionFlag = cmd.VersionCmd.Flags().BoolP("version", "v", true, "api config file (required)")
	cmd.VersionFlag = versionFlag

	rootCmd.AddCommand(startCmd)
	apiConfigFile = startCmd.Flags().StringP("config", "c", "", "api config file (required)")
	startCmd.MarkFlagRequired("config")

}
