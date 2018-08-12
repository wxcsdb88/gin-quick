package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// LoadConfig load the config and output
// according to your input output structure
// wx-kits

//555team

//666team

func LoadConfig(file string, output interface{}) error {
	if "" == file {
		return fmt.Errorf("blank file")
	}

	v := viper.New()
	// use viper auto detect the file(suffix)
	v.SetConfigFile(file)

	err := v.ReadInConfig()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	err = v.Unmarshal(&output)
	return err
}
