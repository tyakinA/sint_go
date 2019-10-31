package sintgo

import "fmt"

func main() {
	var a int
	fmt.Print("a=")
	fmt.Scanln(&a)

	if a%2 == 0 {
		fmt.Println("Четное число")
	} else {
		fmt.Println("Нечётное число")
	}
}
