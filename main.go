package main

import "go_hello_world/src"

func main() {
	println("fire!")

	c := kakezan()
	println(c)

	hwStr := src.Hello()

	println(hwStr)

	src.FileWrite("message.txt", hwStr)
}

func kakezan() int {
	var a int = 6
	var b int = 7
	var c int = a * b
	return c
}
