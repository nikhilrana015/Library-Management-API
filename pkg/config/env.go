package config

import (
	"log"

	"github.com/spf13/viper"
)

var EnvConfigs *envConfigs

type envConfigs struct {
	LocalServerPort string `mapstructure:"LOCAL_SERVER_PORT"`
	DbUsername      string `mapstructure:"DB_USERNAME"`
	DbPassword      string `mapstructure:"DB_PASSWORD"`
	DbServerPort    string `mapstructure:"DB_SERVER_PORT"`
	DbName          string `mapstructure:"DB_NAME"`
}

func loadEnvVariables() (config *envConfigs) {

	// Tells viper the path of your env file. Root path contains as . only
	viper.AddConfigPath("./")
	// name of the env file
	viper.SetConfigName("process")
	// tells viper the type of file
	viper.SetConfigType("env")
	// viper reads all variables from env file and log error if found any
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading environment file", err)
	}
	// viper unmarshals the loaded env variables into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return config
}

func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}
