package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/TranThang-2804/auto-logwork/pkg/constant"
)

func CheckConfigExist() (bool, error) {
	if _, err := os.Stat(constant.ConfigFile); err == nil {
		return true, err
	} else {
		return false, err
	}
}

func ReadConfig() (error) {
  configFileExist, err := CheckConfigExist()

  if err != nil {
    log.Fatal(err)
    return err
  }

  if (configFileExist) {
    fileByte, err := os.ReadFile(constant.ConfigFile)

    if err != nil {
      return err
    }
    
    fmt.Printf("fileByte: %v\n", fileByte)
  }
  
  return nil
}
