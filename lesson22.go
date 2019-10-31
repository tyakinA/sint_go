package sintgo

import "fmt"

func main() {
	var a int
	fmt.Print("a=")
	fmt.Scanln(&a)

	if a%3 == 0 {
		fmt.Println("Число делится на 3")
	} else {
		fmt.Println("Число не делтся на 3")
	}
}
