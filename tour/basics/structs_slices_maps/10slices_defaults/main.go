package main

import "fmt"

// 在进行切片时，你可以利用它的默认行为来忽略上下界。

// 切片下界的默认值为 0，上界则是该切片的长度。
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[0:4]
	fmt.Println(s)

	s = s[:5]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
