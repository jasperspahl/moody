/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package mood

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jasperspahl/moody/internal/model"
	"github.com/jasperspahl/moody/internal/utils"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var name string
var value float64
var colorStr string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a Mood",
	RunE:  addMood,
}

func addMood(cmd *cobra.Command, args []string) (err error) {
	if name == "" {
		namePrompt := promptui.Prompt{
			Label: "Input Name",
		}
		name, err = namePrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}
	}
	if value == 0 {
		valueValidation := func(input string) error {
			val, err := strconv.ParseFloat(input, 64)
			if err != nil || val <= 0 {
				return errors.New("invalid number")
			}
			return nil
		}
		valuePrompt := promptui.Prompt{
			Label:    "Input Value",
			Validate: valueValidation,
		}

		val, err := valuePrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}
		value, _ = strconv.ParseFloat(val, 64)
	}
	if utils.ColorValidation(colorStr) != nil {
		colorPrompt := promptui.Prompt{
			Label:    "Input Color",
			Validate: utils.ColorValidation,
		}

		colorStr, err = colorPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}
	}
	col := utils.ConvertToColor(colorStr)

	created := model.AddMood(name, value, col)
	if !created {
		fmt.Printf("%s is already created\n", name)
	}
	return nil
}

func init() {
	MoodCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the mood")
	addCmd.Flags().Float64VarP(&value, "value", "v", 0, "Value of the mood")
	addCmd.Flags().StringVarP(&colorStr, "color", "c", "", "Color of the mood represented in ARGB")
}
