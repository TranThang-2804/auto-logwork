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

	"github.com/TranThang-2804/auto-logwork/cmd/internal/configure"
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
	configFileExist, _ := configure.CheckConfigExist()

	reader := bufio.NewReader(os.Stdin)

	if configFileExist {
		fmt.Print("Configuration Exists, Overwrite? [y/n]: ")
		overwrite, _ := reader.ReadString('\n')
		overwrite = overwrite[:len(overwrite)-1]

		if overwrite == "n" {
			return nil
		} else if overwrite != "y" {
			log.Println("Invalid input")
			return errors.New("Invalid input, valid input are y/n")
		}
	}

	fmt.Print("Enter type: ")
	endpointType, _ := reader.ReadString('\n')
	endpointType = endpointType[:len(endpointType)-1]
	fmt.Print("Enter endpoint: ")
	endpoint, _ := reader.ReadString('\n')
	endpoint = endpoint[:len(endpoint)-1]
	fmt.Print("Enter credentials: ")
	credential, _ := reader.ReadString('\n')
	credential = credential[:len(credential)-1]

	config := &types.Config{
		EndpointType: endpointType,
		Endpoint:     endpoint,
		Credential:   credential,
	}

	err := configure.WriteConfig(config)

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
