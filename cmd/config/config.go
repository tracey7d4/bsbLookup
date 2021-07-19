package config

import (
	"fmt"
	"github.com/spf13/viper"
)
type Config struct{
	Port	int `mapstructure:"port"`
}

func LoadConfig() (config Config, err error) {
 	viper.SetConfigFile("config/config.yaml")
 	viper.AutomaticEnv()

 	err = viper.ReadInConfig()
 	if err != nil {
 		fmt.Printf("Error reading config file %s\n", err)
 		return
	}
	err = viper.Unmarshal(&config)
	return
}


