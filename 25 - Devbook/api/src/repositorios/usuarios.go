package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// Usuarios representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// Cria um repositório de usuários com o banco de dados fornecido
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados e retorna o ID gerado
func (repositorio *Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statment, erro := repositorio.db.Prepare(
		"INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statment.Close()

	resultado, erro := statment.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}
	UltimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(UltimoIDInserido), nil
}

func (repositorio *Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	statment, erro := repositorio.db.Prepare(
		"SELECT id, nome, nick, email,criadoEm,atualizadoEm FROM usuarios WHERE nome LIKE ? OR nick LIKE ?",
	)
	if erro != nil {
		return nil, erro
	}
	defer statment.Close()

	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
	resultado, erro := statment.Query(nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}
	defer resultado.Close()

	var usuarios []modelos.Usuario

	for resultado.Next() {
		var usuario modelos.Usuario
		if erro = resultado.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
			&usuario.AtualizadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}
