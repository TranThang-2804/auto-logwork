/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/TranThang-2804/auto-logwork/cmd/internal"
	"github.com/TranThang-2804/auto-logwork/pkg/types"
	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "config your credentials and tool's endpoint",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		configureConfig()
	},
}

func configureConfig() error {
	configFileExist, _ := internal.CheckConfigExist()

	reader := bufio.NewReader(os.Stdin)

	if configFileExist {
    fmt.Print("Configuration Exists, Overwrite? [y/n]: ")
    overwrite, _ := reader.ReadString('\n')

    if overwrite == "n" {
      return nil
    } else if overwrite != "y" {
      log.Println("Invalid input")
      return errors.New("Invalid input")
    }
	}

	fmt.Print("Enter type: ")
	endpointType, _ := reader.ReadString('\n')
	fmt.Print("Enter endpoint: ")
	endpoint, _ := reader.ReadString('\n')
	fmt.Print("Enter credentials: ")
	credential, _ := reader.ReadString('\n')

	config := &types.Config{
		EndpointType: endpointType,
		Endpoint:     endpoint,
		Credential:   credential,
	}

	err := internal.WriteConfig(config)

	return err
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
