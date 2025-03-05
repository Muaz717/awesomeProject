package main

import (
	"fmt"
	"sort"
)

type account struct {
	value int
}

func main() {
	//WithoutCopy()
	//WithCopy()

	s1 := make([]account, 0, 2)
	s1 = append(s1, account{})  // len=1, cap=2 [{0}]
	s2 := append(s1, account{}) // len=2, cap=2 [{0} {0}]

	acc := &s2[0]

	acc.value = 100
	fmt.Println(s1, s2) // [{100}] [{100} {0}]

	s1 = append(s2, account{}) // len = 3, cap = 4 [{100} {0} {0}]
	acc.value += 100
	fmt.Println(s1, s2) // [{200} {0}] [{100} {0} {0}]
}

func WithoutCopy() {
	one := []int{1, 2}
	two := one

	two[0] = 123

	fmt.Println(one, two) // [123 2] [123 2]

	one = append(one, 777)

	fmt.Println(one, two) // [123 2 777] [123 2]
}

func WithCopy() {
	one := []int{1, 2, 0, 0, 1}
	two := make([]int, len(one))

	copy(two, one)

	fmt.Println(one, two) // [1 2 0 0 1] [1 2 0 0 1]

	two[0] = 123

	fmt.Println(one, two) // [1 2 0 0 1] [123 2 0 0 1]
}

func sortStruct() {
	var arr = [...]struct{ Name string }{{Name: "b"}, {Name: "c"}, {Name: "a"}}

	fmt.Println(arr) // [{b} {c} {a}]

	sort.SliceStable(arr[:], func(i, j int) bool {
		return arr[i].Name < arr[j].Name
	})

}
