package main

import "fmt"

type User struct {
	Name string
}

func main() {
	a := 25
	fmt.Println(a)

	squareVal(a)
	fmt.Println(a)

	squarePointerVal(&a)
	fmt.Println(a)

	user := User{Name: "Мурад"}
	fmt.Println(user.Name)

	updateUser(&user)
	fmt.Println(user.Name)
}

func updateUser(u *User) {
	u = &User{Name: "Безымянный"}
	fmt.Println("Имя внутри функции [updateUser]:", (*u).Name)

	//u.Name = "Магомед"
}

func squareVal(v int) {
	square := v * v
	v = square
}

func squarePointerVal(v *int) {
	square := *v * *v
	*v = square
}
