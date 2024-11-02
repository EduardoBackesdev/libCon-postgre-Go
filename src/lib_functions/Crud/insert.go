package lib_functions

import (
	"errors"
	"fmt"
	db "lib/src/lib_functions/Connection"
	"strings"
)

// type okStruct struct {
// 	Message string
// 	status  int
// 	Err     Error()
// }

// Funcao principal para query de insert
func insertLib(table string, coluns []string, values []string) {
	db, err := db.OpenConn()
	if err != nil {
		errString := errors.New("Error in connect")
		fmt.Errorf("Error", errString)
		return
	}
	defer db.Close()
	queryString := "INSER INTO" + table + "(" + strings.Join(coluns, ",") + ") VALUES (" + strings.Join(values, ",") + ")"
	_, errQuery := db.Exec(queryString)
	if errQuery != nil {
		fmt.Errorf("Error: ", errQuery)
		return
	}

}

func Insert(t string, c []string, v []string) {
	insertLib(t, c, v)
}
