package main

import "fmt"

func main() {

	array := [5]*int{new(int),new(int)}

		*array[0] =1 
		*array[1] =2



	fmt.Println("Hello, World!", array,)
}
