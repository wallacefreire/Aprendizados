package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	linha, erro := db.Query("SELECT * FROM funcionarios WHERE id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o funcionário!"))
		return
	}

	var funcionario funcionario
	if linha.Next() {
		if erro := linha.Scan(&funcionario.ID, &funcionario.Nome, &funcionario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o funcionário"))
			return
		}
	}

	if erro := json.NewEncoder(w).Encode(funcionario); erro != nil {
		w.Write([]byte("Erro ao converter o funcionário para JSON"))
		return
	}
}

// AtuaizarFuncionario altera os dados de um funcionário no Banco de Dados
func AtualizarFuncionario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var funcionario funcionario
	if erro := json.Unmarshal(corpoRequisicao, &funcionario); erro != nil {
		w.Write([]byte("Erro ao converter o corpo da requisição para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("UPDATE funcionarios SET nome = ?, email = ? WHERE id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(funcionario.Nome, funcionario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o funcionário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
