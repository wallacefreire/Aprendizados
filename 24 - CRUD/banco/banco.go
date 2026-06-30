package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Conectar abre a conexão com o Banco de Dados
func Conectar() (*sql.DB, error) {
	stringConexao := "wallace:wall@/conta_comigo?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
