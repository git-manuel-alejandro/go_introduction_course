package operadores

import "fmt"

func Condicionales(edad int) string {
	if edad >= 18 {
		result := fmt.Sprint("mayor o igual a 18")
		return result

	}

	result := fmt.Sprint("menor a 18")
	return result

}
