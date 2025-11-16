package main

import "fmt"

/*
*
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数数组作为参数，在函数内部将数组中的每个元素增加10，然后在主函数中调用该函数并输出修改后的数组。
考察点 ：指针的使用、值传递与引用传递的区别。
*/
func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("before increase:", nums)
	increase(&nums)
	fmt.Println("after increase:", nums)
}
func increase(nums *[]int) {
	for i := 0; i < len(*nums); i++ {
		(*nums)[i] *= 2
	}
}
