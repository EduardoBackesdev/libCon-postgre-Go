package lib_functions

import (
	"fmt"
	config "lib/src/lib_functions/Connection"
	joinTwo "lib/src/lib_utils"
	"strings"
)

// Funcao principal para query de update
func update(tabela string, colunas []string, values []interface{}) {
	updateValues := []string{}
	db, err := config.OpenConn()
	if err != nil {
		println("Error: ", err)
		panic(err)
	}
	defer db.Close()
	i := joinTwo.JoinLibTwo(values)
	for v, e := range colunas {
		for range i {
			c := ""
			c = e + " = " + i[v]
			updateValues = append(updateValues, c)
			break
		}
	}
	queryString := "UPDATE " + tabela + " set " + strings.Join(updateValues, ", ")
	_, errs := db.Exec(queryString)
	if errs != nil {
		fmt.Println("Error: ", errs)
		return
	}
}

// Funcao de update para atualizar as colunas escolhidas do banco,
// voce precisa passar uma string com o nome da tabela,
// um array de strings com as colunas,
// um array dos valores
// Funcao recebe os seguintes parametros:
//
//	tabela 			string
//	colunas 		[]string
//	values 			[]interface{}
func Update(t string, c []string, v []interface{}) {
	update(t, c, v)
}
