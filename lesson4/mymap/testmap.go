package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]Student)
	fmt.Println(m)
	var s Student
	s.Name = "zhaoyh"
	s.Age = 22
	m["zhaoyh"] = s
	fmt.Println(m["zhaoyh"])

}
