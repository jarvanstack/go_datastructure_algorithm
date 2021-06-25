package main

import "fmt"

type User struct {
	name string
}
func main() {
	str := "hello"
	changeString(str)
	fmt.Printf("name=%#v\n", str)
	user := User{
		name: "hello",
	}
	changeUser(user)
	fmt.Printf("user=%#v\n", user)
}
func changeString(str string)  {
	str = "change"
}
func changeUser(user User)  {
	user.name = "change"
}