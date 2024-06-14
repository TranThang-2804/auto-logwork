/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// logworkCmd represents the logwork command
var logworkCmd = &cobra.Command{
	Use:   "logwork",
	Short: "Auto logwork",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logwork called")
	},
}

func init() {
	rootCmd.AddCommand(logworkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logworkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logworkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
