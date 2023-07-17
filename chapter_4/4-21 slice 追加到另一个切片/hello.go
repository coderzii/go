package main

import "fmt"

// 将一个切片追加到另一个切片
func main(){

	s1 := []int{1,2}
	s2 := []int{3,4}

	res := append(s1,s2...)
	
	fmt.Println(res)
}

