package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Initialize loads the application configuration
func Initialize() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("application")
	viper.AddConfigPath("./dev")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("couldn't read config file. %v\n", err.Error())
	}

	viper.AutomaticEnv()
}
