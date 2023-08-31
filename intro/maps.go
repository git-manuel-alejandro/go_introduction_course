package intro

import (
	"fmt"
	"unsafe"
)

func Maps() {

	m1 := make(map[int]string)
	m1[1] = "manuel1"
	m1[2] = "manuel2"
	m1[3] = "manuel3"
	m1[4] = "manuel4"
	fmt.Println(len(m1))
	fmt.Println(unsafe.Sizeof(m1))

	for i, v := range m1 {
		fmt.Println(i, v)
	}

	m2 := map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}

	fmt.Println(m2)
}
