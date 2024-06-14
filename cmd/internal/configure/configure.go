package configure

import (
	"encoding/json"
	"log"
	"os"

	"github.com/TranThang-2804/auto-logwork/pkg/constant"
	"github.com/TranThang-2804/auto-logwork/pkg/types"
)

func GetConfigFilePath() string {
	homeDir := os.Getenv("HOME")
	return homeDir + "/" + constant.ConfigFile
}

func CheckConfigExist() (bool, error) {
	if _, err := os.Stat(GetConfigFilePath()); err == nil {
		return true, err
	} else {
		return false, err
	}
}

func ReadConfig(config *types.Config) {
	configFileExist, _ := CheckConfigExist()

	if configFileExist {
		file, err := os.Open(GetConfigFilePath())
    defer file.Close()

		if err != nil {
      log.Print(err)
			return
		}
    
    decoder := json.NewDecoder(file)

    err = decoder.Decode(&config)

    if err != nil {
      log.Print(err)
      return
    }

	} else {
		log.Print("You haven't config credentials, to config, run: auto-logwork configure")
    return
	}
}

func WriteConfig(config *types.Config) error {
	file, err := os.Create(GetConfigFilePath())

	if err != nil {
		return err
	}

	defer file.Close()

  encoder := json.NewEncoder(file)
  err = encoder.Encode(config)

	return err
}
