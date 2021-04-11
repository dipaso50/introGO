package main

import "fmt"

func main() {

	//declaración de variables
	//----------------------------------------------------------------------------------------
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	//constantes
	//----------------------------------------------------------------------------------------

	const Pi = 3.14
	fmt.Println(Pi)

	//arrays
	//----------------------------------------------------------------------------------------
	var arr [4]int //Slice sin tamaño definido
	arr[0] = 33
	fmt.Printf("arr --> len(%d) cap(%d)\n", len(arr), cap(arr))

	//slices
	//----------------------------------------------------------------------------------------
	var slcs []int //Slice sin tamaño definido
	slcs = append(slcs, 1, 2, 3)
	fmt.Printf("slcs --> len(%d) cap(%d) %v\n", len(slcs), cap(slcs), slcs)

	var slcs2 []int           //Declaración de slice
	slcs2 = make([]int, 3, 3) //Inicialización con 3 elementos y capacidad de 3
	fmt.Printf("slcs2 --> len(%d) cap(%d)\n", len(slcs2), cap(slcs2))

	//maps
	//----------------------------------------------------------------------------------------

	var m map[string]int     //declaracion
	m = make(map[string]int) //inicialización
	delete(m, "nothing")     // delete a element

	//uso
	m["route"] = 66
	i := m["route"]
	fmt.Printf("m --> len(%d) %d\n", len(m), i)
}
