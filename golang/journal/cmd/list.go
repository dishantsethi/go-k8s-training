package cmd

import (
	"fmt"
	"journal/helper"
	"journal/manager"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all entries of a User",
	Long:  `This command helps user to list entries in the journal.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := manager.ListEntry(Email)
		helper.Check(err)
		fmt.Println("All Entries printed successfully")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&Email, "email", "e", "", "List all entries of the user (required)")
	listCmd.MarkFlagRequired("email")
}
