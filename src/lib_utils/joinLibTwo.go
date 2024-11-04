package libutils

import (
	"fmt"
	"reflect"
)

func JoinLibTwo(a []interface{}) []string {
	newValues := []any{}
	for _, e := range a {
		if reflect.TypeOf(e) == reflect.TypeOf("") {
			a := ""
			a = "'" + e.(string) + "'"
			newValues = append(newValues, a)
		} else {
			newValues = append(newValues, e)
		}
	}
	o := []string{}
	for _, e := range newValues {
		i := fmt.Sprintf("%v", e)
		o = append(o, i)
	}
	return o
}
