package main

import (
	"fmt"
	"math"
	"time"
)

// 记住：方法只是个带接收者参数的函数。
//
// 现在这个 Abs 的写法就是个正常的函数，功能并没有什么变化。
type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
	time.Second * 30
}
