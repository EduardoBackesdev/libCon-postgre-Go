package lib_functions

import (
	"fmt"
	join "lib/lib_utils"
	db "lib/src/lib_functions/Connection"
	"strings"
)

// OBS: VERIFICAR PORQUE NAO DA ERRO NO SQL

// Funcao principal para query de insert
func insertLib(table string, coluns []string, values []interface{}) {
	db, err := db.OpenConn()
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer db.Close()
	queryString := "INSERT INTO " + table + " (" + strings.Join(coluns, ",") + ") VALUES (" + join.JoinLib(values) + ")"
	_, errQuery := db.Exec(queryString)
	if errQuery != nil {
		fmt.Println("Error: ", errQuery)
		panic(errQuery)
	}
}

// Funcao de Insert, voce precisa passar uma string com o nome da tabela,
// um array de strings com as colunas, e um array de strings com os valores
// Funcao os seguintes parametros:
//
//	nome tabela     string
//	colunas     	array[]strings
//	values    		array[]strings
func Insert(t string, c []string, v []interface{}) {
	insertLib(t, c, v)
}
