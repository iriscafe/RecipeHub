package platform

import (
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

type Config struct {
	// Servidor
	Port string
	
	// Banco de dados
	DBDriver string
	DBPath   string
	
	// Logs
	LogLevel string
	
	// Ambiente
	Environment string
}

var AppConfig *Config

func InitConfig() {
	// Carregar arquivo .env se existir
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}
	
	AppConfig = &Config{
        Port: getEnvOrDefault("PORT"),           
        DBDriver: getEnvOrDefault("DB_DRIVER"), 
        DBPath: getEnvOrDefault("DB_PATH"),     
        LogLevel: getEnvOrDefault("LOG_LEVEL"), 
        Environment: getEnvOrDefault("ENVIRONMENT"),
    }
    
    // Validar configurações obrigatórias
    validateConfig()
}

func validateConfig() {
	if AppConfig.Port == "" {
		panic("PORT é obrigatório - defina a variável de ambiente PORT ou crie um arquivo .env")
	}
	if AppConfig.DBDriver == "" {
		panic("DB_DRIVER é obrigatório - defina a variável de ambiente DB_DRIVER ou crie um arquivo .env")
	}
	if AppConfig.DBPath == "" {
		panic("DB_PATH é obrigatório - defina a variável de ambiente DB_PATH ou crie um arquivo .env")
	}
	if AppConfig.LogLevel == "" {
		panic("LOG_LEVEL é obrigatório - defina a variável de ambiente LOG_LEVEL ou crie um arquivo .env")
	}
	if AppConfig.Environment == "" {
		panic("ENVIRONMENT é obrigatório - defina a variável de ambiente ENVIRONMENT ou crie um arquivo .env")
	}
}

func getEnvOrDefault(key string) string {
	return os.Getenv(key)
}

func getEnvOrDefaultInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvOrDefaultBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}