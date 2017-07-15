package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	name string //Json序列化，只能序列化公有对象，不能实例化私有对象
	id   int
}

func (s *Student) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.id)
}

func main() {
	s := &Student{
		name: "binggan",
		id:   1,
	}
	buf, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
}
