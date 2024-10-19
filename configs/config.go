package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Conf struct {
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf

	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("WARNING: %v\n", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
