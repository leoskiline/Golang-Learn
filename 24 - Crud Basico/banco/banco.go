package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de banco de dados MySQL
)

// Conectar abre uma conexão com o banco de dados MySQL e retorna um ponteiro para o objeto sql.DB
// e um erro, caso ocorra. A string de conexão é configurada para se conectar ao banco de dados "devbook"
func Conectar() (*sql.DB, error) {
	stringConexao := "root:golang@tcp(host.docker.internal:3320)/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
