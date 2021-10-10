package main

import (
	"fmt"
)

func updateName(x *string)  {
	*x = "wedge"
}

func main()  {
	name := "tifa"

	m := &name

	// fmt.Println("memeory address of name is: ", m)
	// fmt.Println("value at memory address:", *m)

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)
}