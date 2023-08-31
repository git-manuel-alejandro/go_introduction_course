package tareas

import "fmt"

func Vector() {
	array := [5]int{4, 2, 5, 6, 7}

	for i, v := range array {
		array[i] = v + 20
	}

	fmt.Println("Los valores del array son: ", array)
}
