package lib_functions

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Funcao principal de conexao com o banco de dados
func openConn(host, port, user, password, dbname string) (*sql.DB, error) {
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		errStrings := errors.New("Check the fields and try again!")
		fmt.Errorf(errStrings.Error())
		return nil, errStrings
	}
	db, err := sql.Open("postgres", "host="+host+" port="+port+" user="+user+" password="+password+" dbname="+dbname+" sslmode=disable")
	if err != nil {
		fmt.Errorf(err.Error())
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	return db, err
}

// Funcao de conexao ao banco esperando parametros com as seguintes strings:
//
//	host     string
//	port     string
//	user     string
//	password string
//	dbname   string
func ConnectionString(host, port, user, password, dbname string) (*sql.DB, error) {
	res, err := openConn(host, port, user, password, dbname)
	return res, err
}
