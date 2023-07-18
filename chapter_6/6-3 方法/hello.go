package main

import "fmt"

/*
方法能给用户定义的类型添加新的行为。方法实际上也是函数，只是在声明时，在关键字
func 和方法名之间增加了一个参数
*/

type user struct {
	name string

	email string
}

func main() {
	// user 类型的值可以用来调用
	// 使用值接收者声明的方法
	bill := user{"Bill", "bill@email.com"}

	bill.notify()

	// 指向 user 类型值的指针也可以用来调用
	// 使用值接收者声明的方法

	lisa := &user{"Lisa", "lisa@email.com"}

	// 指向 user 类型值的指针也可以用来调用
	// 使用值接收者声明的方法
	lisa.changeEmail("lisa@newadmin.com")
	lisa.notify()

	// user 类型的值可以用来调用
	// 使用指针接收者声明的方法
	bill.changeEmail("bill@newadmin.com")
	bill.notify()

}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

/*
Go 语言里有两种类型的接收者：值接收者和指针接收者

1.如果使用值接收者声明方法，调用时会使用这个值的一个副本来执行
2.总结一下，值接收者使用\值的副本来调用方法，而指针接受者使用实际值来调用方法


*/
