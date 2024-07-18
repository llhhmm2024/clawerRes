package cmd

import (
	"crawler/bootstrap"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crawler",
	Short: "yzzy crawler of your command",
	Long:  "",
	// rootCmd 的所有子命令都会执行以下代码
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		bootstrap.SetupCfg()
		bootstrap.SetupDB()
		bootstrap.SetupDownload()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
