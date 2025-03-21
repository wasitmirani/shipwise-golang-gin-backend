package config

import (
	// "github.com/spf13/viper"
	// "github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func LoadConfig() Config {
	// if err := godotenv.Load(".env"); err != nil {
	// 	panic("Error loading .env file: " + err.Error())
	// }
	// if _, err := godotenv.Read(".env"); err != nil {
	// 	panic("Error reading .env file: " + err.Error())
	// }
	// viper.AutomaticEnv()

	// if err := viper.ReadInConfig(); err != nil {
	// 	panic("Error reading .env file: " + err.Error())
	// }

	return Config{
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:    "postgres",
		DBPassword: "root",
		DBName:     "nextship_db",
		JWTSecret:  "secret",
	}
}

