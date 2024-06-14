package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/TranThang-2804/auto-logwork/pkg/constant"
	"github.com/TranThang-2804/auto-logwork/pkg/types"
)

func CheckConfigExist() (bool, error) {
	if _, err := os.Stat(constant.ConfigFile); err == nil {
		return true, err
	} else {
		return false, err
	}
}

func GetConfigFilePath() string {
	homeDir := os.Getenv("HOME")
	return homeDir + "/" + constant.ConfigFile
}

func ReadConfig() error {
	configFileExist, _ := CheckConfigExist()

	if configFileExist {
		fileByte, err := os.ReadFile(GetConfigFilePath())

		if err != nil {
			return err
		}

		fmt.Printf("fileByte: %v\n", fileByte)
	} else {
		log.Print("You haven't config credentials, to config, run: auto-logwork configure")
	}

	return nil
}

func WriteConfig(config *types.Config) error {
	file, err := os.Create(GetConfigFilePath())

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Credential: %sEndpoint: %sEndpointType: %s", config.Credential, config.Endpoint, config.EndpointType))

	return err
}
