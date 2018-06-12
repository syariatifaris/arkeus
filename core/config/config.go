package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/syariatifaris/arkeus/core/config/toml"
)

const (
	ExtTomlFile = "toml"
)

func ReadConfigurationModule(moduleName, ext string, dirs []string, cfgContainer interface{}) error {
	filePath := resolveConfigModuleFileLoc(moduleName, ext, dirs)
	if filePath == "" {
		return errors.New("unable to find files path")
	}

	switch ext {
	case ExtTomlFile:
		configuration := toml.NewTomlConfiguration(filePath)
		err := configuration.DecodeConfig(cfgContainer)
		if err != nil {
			return err
		}
	default:
		return errors.New("unrecognized config file extension")
	}

	return nil
}

func resolveConfigModuleFileLoc(module, ext string, dirs []string) string {
	for _, dir := range dirs {
		file := fmt.Sprintf("%s%s.%s", dir, module, ext)
		if _, err := os.Stat(file); !os.IsNotExist(err) {
			return file
		}
	}

	return ""
}
