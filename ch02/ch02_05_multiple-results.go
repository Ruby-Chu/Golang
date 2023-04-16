package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("str1", "str2")
	fmt.Printf("a= %s, b= %s\n", a, b)
}
