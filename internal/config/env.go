package config

import "github.com/spf13/viper"

func Load(path string) (*Config, error) {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
