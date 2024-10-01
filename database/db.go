package database

import (
	"fmt"
	"log"
	"os"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	// Pegando as variáveis de ambiente
	endereco := os.Getenv("DB_HOST")
	usuario := os.Getenv("DB_USER")
	senha := os.Getenv("DB_PASSWORD")
	nomeBanco := os.Getenv("DB_NAME")
	portaBanco := os.Getenv("DB_PORT")

	// Construindo a string de conexão com o banco de dados
	stringDeConexao := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		endereco, usuario, senha, nomeBanco, portaBanco,
	)

	// Abrindo a conexão com o banco de dados
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panicf("Erro ao conectar com banco de dados: %v", err)
	}

	// Migrando o modelo 'Aluno'
	if err = DB.AutoMigrate(&models.Aluno{}); err != nil {
		log.Panicf("Erro ao migrar o modelo: %v", err)
	}

	log.Println("Conexão com banco de dados estabelecida com sucesso")
}
