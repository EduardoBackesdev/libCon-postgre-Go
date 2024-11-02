package lib_functions

import (
	"database/sql"
	"fmt"
	config "lib/src/lib_functions/Connection"
)

// Funcao principal para query de select *
func selectAll(tabela string) (*sql.Rows, error) {
	db, err := config.OpenConn()
	if err != nil {
		println("Error: ", err)
		panic(err)
	}
	defer db.Close()
	queryString := "SELECT * FROM " + tabela
	valor, errQuery := db.Query(queryString)
	if errQuery != nil {
		println("Error", errQuery)
		return nil, errQuery
	}
	return valor, nil
}

// Funcao de select para selecionar todos os itens de uma seguinte tabela,
// voce precisa passar uma string com o nome da tabela,
// Funcao recebe os seguintes parametros:
//
//	tabela 			string
func SelectAll(t string) (*sql.Rows, error) {
	v, e := selectAll(t)
	if e != nil {
		fmt.Println("Error: ", e)
		return nil, e
	}
	return v, nil
}
