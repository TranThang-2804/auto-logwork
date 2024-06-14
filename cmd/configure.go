/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
  "github.com/TranThang-2804/auto-logwork/cmd/internal/configuration"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "config your credentials and tool's endpoint",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configure called")
	},
}

func configureConfig() error {
  configFileExist, err = CheckConfigExist()

  if err != null {
    log.Fatal(err)
    return err
  }

  if (configFileExist) {
    fmt.Println("Configuration Exists, Overwrite? [y/n]")
  }
  
  return nil
}

func init() {
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
