package main

import "fmt"

func main() {

	array := [5]int{1, 2, 3, 4, 5}

	array_1 := [...]int{1, 2, 3}

	array_2 := [5]*int{0: new(int), 1: new(int)}

	*array_2[0] = 10
	*array_2[1] = 20

	fmt.Println("Hello, World!", array, array_1, array_2)
}
