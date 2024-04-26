/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli_cobra/internal/database"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		categoryDB := GetCategory(GetDB())
		cat, err := categoryDB.Create(name, description)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(cat)
		}
	},
}

var name string
var description string

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.Flags().StringVarP(&name, "name", "n", "", "Category name")
	categoryCmd.Flags().StringVarP(&description, "description", "d", "", "Category description")

	categoryCmd.MarkFlagRequired("name")
	categoryCmd.MarkFlagRequired("description")
}

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}

	return db
}

func GetCategory(db *sql.DB) *database.Category {
	return database.NewCategory(db)
}
