package intro

import "fmt"

func ForNormal() {
	sum := 0

	for i := 0; i < 10; i++ {

		sum++
		fmt.Println(i, sum)
	}

	fmt.Println(sum)

}

func ForWhile() {
	i := 0
	for true {
		i++
		fmt.Println(i * i * i * i)
	}
}
