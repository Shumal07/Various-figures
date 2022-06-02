package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	var (
		a0 int = a / 100
		a1 int = (a / 10) % 10
		a2 int = a % 10
	)

	switch {
	case a0 != a1 && a1 != a2 && a0 != a2:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
