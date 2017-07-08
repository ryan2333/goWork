package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		x  int
		x1 int8  = -128 //8表示位数，即1个字节(取值范围：-128-127)
		y  int32        //int32表示在32位系统上的int类型  取值范围：(-2^16 -- 2^16-1)
		z  int64        //int64表示在64位系统上的int类型  取值范围：(-2^32 -- 2^32-1)
		u  uint         //u表示无符号整数
		u1 uint8 = 255  //u表示无符号整数；8表示位数，即1个字节(取值范围：0-255)
		v  uint32
		w  uint64
	)
	fmt.Println("x int= ", x, "byte=", unsafe.Sizeof(x))
	fmt.Println("x int32=", y, "byte=", unsafe.Sizeof(y))
	fmt.Println("x int64=", z, "byte=", unsafe.Sizeof(z))
	fmt.Println("x uint=", u, "byte=", unsafe.Sizeof(u))
	fmt.Println("x uint32=", v, "byte=", unsafe.Sizeof(v))
	fmt.Println("x uint64=", w, "byte=", unsafe.Sizeof(w))
	fmt.Println("x1 int8= ", x1, "byte=", unsafe.Sizeof(x1))
	fmt.Println("u1 uint8= ", u1, "byte=", unsafe.Sizeof(u1))
}
