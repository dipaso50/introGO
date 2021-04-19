package main

import (
	"fmt"
	"strings"
)

var separator = "-"

func concat(values ...interface{}) (string, error) {
	strs := make([]string, len(values))
	for i, v := range values {
		fmt.Println(v)
		if val, ok := v.(string); ok && len(val) == 0 {
			return "", fmt.Errorf("Invalid parameter %v", val)
		}
		strs[i] = fmt.Sprintf("%v", v)
	}
	return strings.Join(strs, separator), nil
}

func concatString(params ...string) (string, error) {
	return concat(params[0:])
}

func concatInt(params ...int) (res string) {
	res, _ = concat(params[0:])
	return
}

func main() {
	val := concatInt(3, 5, 7)

	val, err := concatString("hola", "adios")

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(val)
	}

}
