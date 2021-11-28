package main

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}
type Account struct {
	Id int
	//Name    string
	Cleaner func(string) string
	Owner   Person
	Person
}

func main() {
	var a int = 5
	fmt.Println(a)
	for i := 0; i < 2; i++ {
		fmt.Println("loop iteration", i)
		if i == 1 {
			continue
		}
	}
	// complex64, complex128
	var c complex128 = -1.1 + 7.12i
	c2 := -1.1 + 7.12i

	fmt.Println(c, c2)
}
