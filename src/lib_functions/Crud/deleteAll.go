package lib_functions

import (
	config "lib/src/lib_functions/Connection"
)

// Funcao principal para query de select *
func deleteAll(tabela string) {
	db, err := config.OpenConn()
	if err != nil {
		println("Error: ", err)
		panic(err)
	}
	stringQuery := "DELETE FROM " + tabela
	_, errQuery := db.Exec(stringQuery)
	if errQuery != nil {
		println("Error: ", err)
		panic(err)
	}
}

// Funcao de delete para deletar todos os dados de uma tabela,
// CUIDADO!!!
// Funcao recebe os seguintes parametros:
//
//	tabela 			string
func DeleteAll(t string) {
	deleteAll(t)
}
