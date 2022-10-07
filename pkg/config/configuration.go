package config

import "github.com/spf13/viper"

var Config Configuraiton

type Configuraiton struct {
	Database DatabaseConfig
	Server   ServerConfig
}

func Init() error {
	var configuration Configuraiton

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		//log
		return err
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		//log
		return err
	}

	Config = configuration
	return nil
}

func GetGlobalConfig() Configuraiton {
	return Config
}
