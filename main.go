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
	var a int = 25
	fmt.Println(a)
}
