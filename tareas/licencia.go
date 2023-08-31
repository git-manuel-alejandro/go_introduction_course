package tareas

import "fmt"

func Licencia(license bool, age int) {
	if age > 15 && license == true {
		fmt.Println("Puede seguir avanzando")
		return
	}
	fmt.Println("No puede circular")
}
