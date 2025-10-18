package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	viper.AutomaticEnv()
}

func GetDSN() string {
	return viper.GetString("DB_USER") + ":" +
		viper.GetString("DB_PASS") + "@tcp(" +
		viper.GetString("DB_HOST") + ":" +
		viper.GetString("DB_PORT") + ")/" +
		viper.GetString("DB_NAME") + "?charset=utf8mb4&parseTime=True"
}
