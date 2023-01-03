// Package cmd
/*
Copyright © 2023 Jasper Spahl <jasperspahl@web.de>
*/
package cmd

import (
	"fmt"

	"github.com/jasperspahl/moody/internal/model"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the database",
	Long: `Initialize the database with some initial moods and activities.

The moods which will be created will look as followed:
{
	"name": "Super",
	"value": 5.0,
},
{
	"name": "Good",
	"value": 4.0,
},
{
	"name": "Ok",
	"value": 3.0,
},
{
	"name": "Bad",
	"value": 2.0,
},
{
	"name": "Ugly",
	"value": 1.0,
}
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Seeding moods")
		model.SeedMoods()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
