package lib_functions

import (
	"fmt"
	config "lib/src/lib_functions/Connection"
	joinTwo "lib/src/lib_utils"
	"strings"
)

// VERIFICAR ERRO NOS FOR

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
	for _, e := range colunas {
		for _, n := range i {
			c := ""
			c = e + " = " + n
			updateValues = append(updateValues, c)
		}
	}
	queryString := "UPDATE " + tabela + " set " + strings.Join(updateValues, ",")
	fmt.Println(queryString)
	// _, errs := db.Exec(queryString)
	// if errs != nil {
	// 	fmt.Println("Error: ", errs)
	// 	return
	// }
}

// Funcao de select para selecionar todos os itens de uma seguinte tabela,
// voce precisa passar uma string com o nome da tabela,
// Funcao recebe os seguintes parametros:
//
//	tabela 			string
func Update(t string, c []string, v []interface{}) {
	update(t, c, v)
}
