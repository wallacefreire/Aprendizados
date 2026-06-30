package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type funcionario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarFuncionario dá acesso ao método POST
func CriarFuncionario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler"))
		return
	}

	var funcionario funcionario

	if erro = json.Unmarshal(corpoRequisicao, &funcionario); erro != nil {
		w.Write([]byte("Erro ao converter funcionário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro e mais erro"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into funcionarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(funcionario.Nome, funcionario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o ID inserido!! :("))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Funcionário inserido com sucesso! ID: %d", idInserido)))
}

// BuscarFuncionario traz todos os funcionários salvos no Banco de Dados
func BuscarFuncionarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM funcionarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar"))
		return
	}
	defer linhas.Close()

	var funcionarios []funcionario
	for linhas.Next() {
		var funcionario funcionario

		if erro := linhas.Scan(&funcionario.ID, &funcionario.Nome, &funcionario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o funcionário"))
			return
		}

		funcionarios = append(funcionarios, funcionario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(funcionarios); erro != nil {
		w.Write([]byte("Erro ao converter os funcionarios para JSON"))
		return
	}
}

// BuscarFuncionario traz um funcionário específico do Banco de Dados
func BuscarFuncionario(w http.ResponseWriter, r *http.Request) {

}
