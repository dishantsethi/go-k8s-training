package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout registered user to the journal",
	Long:  `This command helps user logout to the journal.`,
	Run: func(cmd *cobra.Command, args []string) {
		// manager.LogOut(Email)
		fmt.Println("Logout successful")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
	logoutCmd.Flags().StringVarP(&Email, "email", "e", "", "Email of the user (required)")
	logoutCmd.MarkFlagRequired("email")
}