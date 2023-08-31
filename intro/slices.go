package intro

import "fmt"

func Slices() {

	var slices []int

	for i := 0; i < 9; i++ {
		slices = append(slices, i)

	}
	var slices1 []int = slices

	fmt.Println(slices)
	fmt.Println(slices1)

	slices[0] = 100

	fmt.Println(slices)
	fmt.Println(slices1)

}
