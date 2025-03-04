package main

import "fmt"

func main() {
	a := 25
	fmt.Println(a)

	squareVal(a)
	fmt.Println(a)

	squarePointerVal(&a)
	fmt.Println(a)

}

func squareVal(v int) {
	square := v * v
	v = square
}

func squarePointerVal(v *int) {
	square := *v * *v
	*v = square
}
