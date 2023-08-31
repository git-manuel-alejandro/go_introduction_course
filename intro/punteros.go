package intro

import (
	"fmt"
)

var numero int
var direccion = &numero

func Punteros() {
	numero = 369

	// & se hace referencia a la direccion

	// * se hace referencia al valor que tiene determinada dirección
	*direccion = numero + 1
	fmt.Println(numero, *direccion)

}

func Add(n *int) {
	fmt.Println("fn add", n)

	*n = *n + 1

}

type MyStruc struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func SlicesPuntero() {
	var slice []int

	for i := 0; i < 9; i++ {
		slice = append(slice, i)
	}

	var slice2 = slice

	slice2[0] = 99

	var direccion = &slice[1]

	*direccion = 100

	fmt.Println(&slice[0])
	fmt.Println(&slice2[0])
	slice2 = append(slice2, 100)
	fmt.Println(slice)
	fmt.Println(slice2)
}

func PunteroSlice() {
	var p1 = MyStruc{ID: 1, Name: "manuel"}
	fmt.Println(p1)

}
