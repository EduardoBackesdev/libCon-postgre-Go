package lib_functions

import (
	"errors"
	"fmt"
	db "lib/src/lib_functions/Connection"
	"strings"
)

// OBS: VERIFICAR PORQUE NAO DA ERRO NO SQL

// Funcao principal para query de insert
func insertLib(table string, coluns []string, values []string) {
	newValues := []string{}
	for _, e := range values {
		fmt.Println(e)
		a := ""
		a = "'" + e + "'"
		newValues = append(newValues, a)
	}
	db, err := db.OpenConn()
	if err != nil {
		errString := errors.New("Error in connect")
		fmt.Errorf("Error", errString)
		return
	}
	defer db.Close()
	queryString := "INSERT INTO " + table + " (" + strings.Join(coluns, ",") + ") VALUES (" + strings.Join(newValues, ",") + ")"
	fmt.Println(queryString)
	_, errQuery := db.Exec(queryString)
	if errQuery != nil {
		fmt.Errorf("Error: ", errQuery)
		return
	}
}

// Funcao de Insert, voce precisa passar uma string com o nome da tabela,
// um array de strings com as colunas, e um array de strings com os valores
// Funcao os seguintes parametros:
//
//	nome tabela     string
//	colunas     	array[]strings
//	values    		array[]strings
func Insert(t string, c []string, v []string) {
	insertLib(t, c, v)
}
