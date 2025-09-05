package platform

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	
	// Usar configuração do banco
	dbPath := AppConfig.DBPath
	LogInfo("Conectando ao banco: " + dbPath)
	
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		LogError("Falha ao conectar com o banco de dados: " + err.Error())
		panic("Falha ao conectar com o banco de dados: " + err.Error())
	}
	
	LogInfo("Banco de dados conectado com sucesso!")
}