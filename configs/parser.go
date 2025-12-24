package configs

import (
	"os"
	"fmt"
	"encoding/json"
)

func parseAppConfigs(appConfigFile string, appConfigs *AppConfigs) error {
	appConfigFileBytes, err := os.ReadFile(appConfigFile)
	if err != nil {
		fmt.Printf("Failed to read from file '%s'\n", appConfigFile)
		return err
	}

	err = json.Unmarshal(appConfigFileBytes, appConfigs)
	if err != nil {
		fmt.Printf("Failed to json unmarshal the bytes of file '%s'\n", appConfigFile)
		return err
	}

	return nil
} 
