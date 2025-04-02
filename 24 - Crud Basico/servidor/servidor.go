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

type usuario struct {
	Id   uint32 `json:"id"`
	Nome string `json:"nome"`
}

// CriarUsuario insere um novo usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpo, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpo, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o corpo da requisição em JSON"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO usuarios (nome) VALUES (?)")
	if erro != nil {
		w.Write([]byte("Erro ao preparar a declaração SQL"))
		return
	}
	defer statement.Close()
	insercao, erro := statement.Exec(usuario.Nome)
	if erro != nil {
		w.Write([]byte("Erro ao executar a inserção"))
		return
	}
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o ID do usuário inserido"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))
}

// Buscar todos os usuários
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()
	linhas, erro := db.Query("SELECT id, nome FROM usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuários"))
		return
	}
	defer linhas.Close()
	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario
		if erro = linhas.Scan(&usuario.Id, &usuario.Nome); erro != nil {
			w.Write([]byte("Erro ao escanear os dados do usuário"))
			return
		}
		usuarios = append(usuarios, usuario)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if erro = json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuários em JSON"))
		return
	}
}

// Buscar usuário por ID
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o ID do usuário"))
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()
	linhas, erro := db.Query("SELECT id, nome FROM usuarios WHERE id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuário"))
		return
	}
	defer linhas.Close()
	var usuario usuario
	if linhas.Next() {
		if erro = linhas.Scan(&usuario.Id, &usuario.Nome); erro != nil {
			w.Write([]byte("Erro ao escanear os dados do usuário"))
			return
		}
	} else {
		w.Write([]byte("Usuário não encontrado"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if erro = json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário em JSON"))
		return
	}
}

// Atualizar usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o ID do usuário"))
		return
	}
	corpo, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}
	var usuario usuario
	if erro = json.Unmarshal(corpo, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o corpo da requisição em JSON"))
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()
	statement, erro := db.Prepare("UPDATE usuarios SET nome = ? WHERE id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao preparar a declaração SQL"))
		return
	}
	defer statement.Close()
	_, erro = statement.Exec(usuario.Nome, ID)
	if erro != nil {
		w.Write([]byte("Erro ao executar a atualização"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Deletar usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o ID do usuário"))
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()
	statement, erro := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao preparar a declaração SQL"))
		return
	}
	defer statement.Close()
	_, erro = statement.Exec(ID)
	if erro != nil {
		w.Write([]byte("Erro ao executar a exclusão"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
