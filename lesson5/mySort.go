package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Id   int
}

func main() {
	s1 := []int{2, 3, 4, 1, 5, 9, 7}
	// s1 := []rune("hello")
	sort.Slice(s1, func(i, j int) bool { //第一个参数为要排序的，匿名函数返回True or False
		//	return s1[i] < s1[j] //如果下标为i的元素小于下标为j的元素则返回True(从小到大排序)
		return s1[i] > s1[j] //如果下标为i的元素大于下标为j的元素则返回True(从大到小排序)
	})
	// fmt.Println(string(s1))
	fmt.Println(s1)

	ss := []Student{}
	ss = append(ss, Student{
		Name: "aa",
		Id:   2,
	})
	ss = append(ss, Student{
		Name: "bb",
		Id:   1,
	})
	ss = append(ss, Student{
		Name: "dd",
		Id:   4,
	})
	ss = append(ss, Student{
		Name: "dd",
		Id:   3,
	})
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Id < ss[j].Id //以ID从小到大排序
	})
	fmt.Println(ss)

}
