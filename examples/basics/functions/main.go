package main

import "fmt"

func concatParams(param1, param2 string, param3 int) (string, error) {
	if len(param1) == 0 || len(param2) == 0 {
		return "", fmt.Errorf("Invalid parameters")
	}

	return fmt.Sprint(param1, "-", param2, "-", param3), nil
}

func concatParamsNamed(param1, param2 string, param3 int) (res string, err error) {
	if len(param1) == 0 || len(param2) == 0 {
		return "", fmt.Errorf("Invalid parameters")
	}

	res = fmt.Sprint(param1, "-", param2, "-", param3)
	return
}

func main() {
	val, err := concatParams("hola", "", 1)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(val)
	}

	val, err = concatParamsNamed("hola", "mundo", 2)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(val)
	}

}
