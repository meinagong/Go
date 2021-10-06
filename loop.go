package main

import (
	"fmt"
)

func main()  {
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	for i := 0; i < 4; i++ {
		fmt.Println("value of i is:", i)
	}

	names := []string{"mario", "meina", "sarah"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("the name at index %v is %v \n", index, value)
	}

	for _, value := range names {
		fmt.Printf("the name is %v \n", value)
	}

	for _, value := range names {
		fmt.Printf("the name is %v \n", value)
		value = "new string" // not work, can't update the value
	}

	fmt.Println(names)
}