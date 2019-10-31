package sintgo

import "fmt"

func fib(N uint) uint {
	var a []uint
	a = make([]uint, N+1)
	a[0] = 0
	a[1] = 1

	for i := uint(2); i <= N; i += 1 {
		a[i] = a[i-1] + a[i-2]
	}

	return a[N]
}
func main() {
	fmt.Println(fib(100))
}
