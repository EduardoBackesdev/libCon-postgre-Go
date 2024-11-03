package lib_functions

import (
	"errors"
	"fmt"
	db "lib/src/lib_functions/Connection"
	"reflect"
	"strings"
)

// OBS: VERIFICAR PORQUE NAO DA ERRO NO SQL, ADICIONAR INSER COM OS DADOS DO BANCO SE NAO VAI DA B.O

// Funcao principal para query de insert
func insertLib(table string, coluns []string, values []interface{}) {
	newValues := []any{}
	for _, e := range values {
		if reflect.TypeOf(e) == reflect.TypeOf("") {
			a := ""
			a = "'" + e.(string) + "'"
			newValues = append(newValues, a)
		} else {
			newValues = append(newValues, e)
		}
	}
	fmt.Println(strings.Join(newValues))
	db, err := db.OpenConn()
	if err != nil {
		errString := errors.New("Error in connect")
		fmt.Errorf("Error", errString)
		return
	}
	defer db.Close()
	// queryString := "INSERT INTO " + table + " (" + strings.Join(coluns, ",") + ") VALUES (" + strings.Join(newValues, ",") + ")"
	// fmt.Println(queryString)
	// _, errQuery := db.Exec(queryString)
	// if errQuery != nil {
	// 	fmt.Errorf("Error: ", errQuery)
	// 	return
	// }
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
