package main

import (
	"fmt"
)

func main()  {
	menu := map[string]float64 {
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["salad"])

	// looping maps
	for k,v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string {
		267684967: "Meina",
		984759373: "Alex",
		845775485: "Sarah",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[845775485])

	phonebook[984759373] = "bowser"
	fmt.Println(phonebook)
}