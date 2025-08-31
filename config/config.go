package config

import "os"

// Config guarda la información del puerto y el host de la aplicacion.
type Config struct {
	Host string // Dirección del servidor
	Port string // Puerto del servidor
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
	return Config{
		Host: getEnv("HOST", "0.0.0.0"),
		Port: getEnv("PORT", "8080"),
	}
}
