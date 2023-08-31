package intro

import (
	"fmt"
	"unsafe"
)

func Vectores() {

	// son de largo fijo
	var vector [6]int

	// for i := 0; i < len(vector); i++ {
	// 	vector[i] = i
	// }

	for i := range vector {
		vector[i] = (i + 10000000) * 2000
	}

	fmt.Println(vector)
	fmt.Println(unsafe.Sizeof(vector))
}
