package main

import (
	config "lib/src/lib_functions/Connection"
)

func main() {
	config.ConnectionConfig("localhost", "5432", "postgres", "admin", "tester")
}
