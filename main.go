package main

import (
	config "lib/src/lib_functions/Connection"
	insert "lib/src/lib_functions/Crud"
)

func main() {
	config.ConnectionConfig("localhost", "5432", "postgres", "admin", "tester")
	insert.Insert("tabela_teste", []string{"nome"}, []string{"Eduardo"})
}
