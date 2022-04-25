package main

import "fmt"

// 每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。

// 类型 []T 表示一个元素类型为 T 的切片。

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// 以下表达式创建了一个切片，它包含 a 中下标从 1 到 3 的元素：
	var s []int = primes[1:4]
	fmt.Println(s)
}
