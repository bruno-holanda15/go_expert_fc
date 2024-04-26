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
	Short: "Create Category.",
	Long: `Create Category new version.`,
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
}
