package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// 结构体字段使用点号来访问。
func main() {
	v := Vertex{1, 2}
	fmt.Println(&(v.X))
	v.X = 4
	fmt.Println(&(v.X))
	fmt.Println(v.X)
}
