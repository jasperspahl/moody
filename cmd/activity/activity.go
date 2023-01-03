/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package activity

import (
	"github.com/spf13/cobra"
)

// activityCmd represents the activity command
var ActivityCmd = &cobra.Command{
	Use:   "activity",
	Short: "Manage Activities",
	Long:  `Work with Activities`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// activityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// activityCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
