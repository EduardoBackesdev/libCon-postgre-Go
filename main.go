package main

import (
	config "lib/src/lib_functions/Connection"
	crud "lib/src/lib_functions/Crud"
)

func main() {
	var a []interface{}
	a = []interface{}{"eduardo", 2}
	config.ConnectionConfig("localhost", "5432", "postgres", "ADMIN", "tester")
	crud.Insert("tabela_teste", []string{"nome"}, a)

}
