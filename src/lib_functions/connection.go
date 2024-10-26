package lib_functions

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Struct para retornar dados do tipo referente ao banco
type sdata struct {
	res *sql.DB
	err error
}

// Struct para parametros da funcao de conexao do banco
type sConnection struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

// Funcao principal de conexao com o banco de dados
func openConn(host, port, user, password, dbname string) (*sql.DB, error) {
	db, err := sql.Open("postgres", "host="+host+" port="+port+" user="+user+" password="+password+" dbname="+dbname+" sslmode=disable")
	if err != nil {
		panic(err) // TRATAR MELHOR O ERRO
	}
	err = db.Ping()
	return db, err
}

// Funcao exportada para lib aceitando Struct
func ConnectionStruct(params sConnection) sdata {
	res, err := openConn(params.host, params.port, params.user, params.password, params.dbname)
	return sdata{res: res, err: err}
}

// Funcao exportada para lib aceitando Strings
func ConnectionString(host, port, user, password, dbname string) sdata {
	res, err := openConn(host, port, user, password, dbname)
	return sdata{res: res, err: err}
}
