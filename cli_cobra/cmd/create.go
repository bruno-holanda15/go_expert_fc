/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli_cobra/internal/database"

	"github.com/spf13/cobra"
)

func newCmdCreateCategory(catDB *database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		_, err := catDB.Create(name, description)
		if err != nil {
			return err
		}

		return nil
	}
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	RunE: newCmdCreateCategory(GetCategory(GetDB())),
}

var name string
var description string

func init() {
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&name, "name", "n", "", "Category name")
	createCmd.Flags().StringVarP(&description, "description", "d", "", "Category description")

	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("description")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
