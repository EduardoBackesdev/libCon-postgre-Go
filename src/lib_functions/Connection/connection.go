package lib_functions

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type dataDbStruct struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

var DataDb dataDbStruct

// Funcao de configuracao do banco de dados, recomendado
// deixar esta funcao em arquivo global de config do seu sistema
// pois para utilizar outras funcoes esta precisa estar configurada
// Funcao espera parametros com as seguintes strings:
//
//	host     string
//	port     string
//	user     string
//	password string
//	dbname   string
func ConnectionConfig(host, port, user, password, dbname string) (string, error) {
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		errStrings := errors.New("Check the fields and try again!")
		fmt.Println("Error: ", errStrings)
		return "", errStrings
	}
	DataDb = dataDbStruct{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname,
	}
	return "OK", nil
}

// Funcao principal de conexao com o banco de dados
func OpenConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host="+DataDb.host+" port="+DataDb.port+" user="+DataDb.user+" password="+DataDb.password+" dbname="+DataDb.dbname+" sslmode=disable")
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	return db, err
}
