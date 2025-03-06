package main

import (
	"fmt"
)

type User struct {
	Name string
}

func main() {
	user := User{Name: "Олег"}
	fmt.Println("Имя до обновления:", user.Name) // Олег

	updateUser(user)
	fmt.Println("Имя после обновления:", user.Name) // Олег

	resetUser(&user)
	fmt.Println("Имя после вызова функции [resetUser] внутри функции [updateUser]:", user.Name) // Безымянный
}

func updateUser(u User) {
	u.Name = "Таненбаум"
	fmt.Println("Имя внутри функции [updateUser]:", u.Name) // Таненбаум

	resetUser(&u)
	fmt.Println("Имя после вызова функции [resetUser] внутри функции [updateUser]:", u.Name) // Безымянный
}

func resetUser(u *User) {
	u = &User{Name: "Безымянный"}
	fmt.Println("Имя внутри функции [resetUser]:", (*u).Name) // Безымянный  Безымянный
}
