// package config

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// func Connect() *gorm.DB {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 		os.Getenv("DB_USER"),
// 		os.Getenv("DB_PASS"),
// 		os.Getenv("DB_HOST"),
// 		os.Getenv("DB_PORT"),
// 		os.Getenv("DB_NAME"),
// 	)

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return db
// }

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
