package config

import (
	"fmt"
	"os"
)

// Config guarda la información de entorno.
type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DSN        string
}

// getEnv obtiene el valor de una variable de entorno dada su clave.
// Si no existe, retorna el parametro proporcionado. Utiliza os.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// LoadConfig retorna un struct Config con el puerto y el host. Por defecto,
// si no se usan parametros, es localhost:8080.
func LoadConfig() Config {
	config := Config{}
	config.AppPort = getEnv("APP_PORT", "8080")

	config.DBHost = getEnv("POSTGRES_HOST", "localhost")
	config.DBPort = getEnv("POSTGRES_PORT", "5432")
	config.DBUser = getEnv("POSTGRES_USER", "admin")
	config.DBPassword = getEnv("POSTGRES_PASSWORD", "password")
	config.DBName = getEnv("POSTGRES_DB", "movies_db")

	config.DSN = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName)

	return config
}
