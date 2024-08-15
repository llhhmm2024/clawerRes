package cmd

import (
	"crawler/bootstrap"

	"github.com/spf13/cobra"
)

var AutoCmd = &cobra.Command{
	Use:   "manual",
	Short: "手动同步任务",
	Long:  `手动同步`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.SetupTask()
	},
}
var CronAutoCmd = &cobra.Command{
	Use:   "cron",
	Short: "定时任务同步数据",
	Long:  `定时任务同步数据`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.RunCronTask()
	},
}

var SpecialCmd = &cobra.Command{
	Use:   "special",
	Short: "指定 video ID 同步",
	Long:  `定时任务同步数据`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.RunSpecialTask()
	},
}

var updateM3u8Cmd = &cobra.Command{
	Use:   "m3u8",
	Short: "Synchronize data from scheduled tasks by special",
	Long:  `定时任务同步数据`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.UpdateM3u8()
	},
}

func init() {
	rootCmd.AddCommand(AutoCmd)
	rootCmd.AddCommand(CronAutoCmd)
	rootCmd.AddCommand(SpecialCmd)
	rootCmd.AddCommand(updateM3u8Cmd)
}
