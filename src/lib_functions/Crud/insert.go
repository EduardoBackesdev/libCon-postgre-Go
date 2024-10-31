package lib_functions

import (
	"errors"
	"fmt"
	db "lib/src/lib_functions/Connection"
)

// Funcao principal para query de insert
func insertLib(table string, coluns []string, values []string) {
	db, err := db.OpenConn()
	if err != nil {
		errString := errors.New("Error in connect!")
		fmt.Errorf("Error", errString)
		return
	}
	defer db.Close()
	queryString := ""
	_, errQuery := db.Exec()
	if errQuery != nil {
		fmt.Errorf("Error: ", errQuery)
		return
	}

}

func Insert() {

}
