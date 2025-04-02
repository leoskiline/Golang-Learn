package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "root:golang@tcp(host.docker.internal:3320)/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro.Error())
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro.Error())
	}

	fmt.Println("Conexão está aberta!")

	rows, erro := db.Query("SELECT id, nome FROM usuarios")
	if erro != nil {
		log.Fatal(erro.Error())
	}
	defer rows.Close()

	fmt.Println(rows)
}
