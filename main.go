package main

import (
	config "lib/src/lib_functions/Connection"
	crud "lib/src/lib_functions/Crud"
)

func main() {
	var a []interface{}
	a = []interface{}{"Andriele", 25, "Oliver"}
	config.ConnectionConfig("localhost", "5432", "postgres", "admin", "tester")
	crud.Update("tabela_teste", []string{"nome", "idade", "cachorro"}, a)

}
