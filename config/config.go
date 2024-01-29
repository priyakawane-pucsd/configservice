package config

import (
	"configservice/repository/mongo"

	"github.com/spf13/viper"
)

// Config holds the configuration parameters.
type Config struct {
	Name  string
	Build string // dev,local,prod
	Mongo mongo.Config
}

// LoadConfig loads the configuration from file.
func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(err)
	}

	return config
}
