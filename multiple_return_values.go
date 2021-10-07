package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main()  {
	firstName1, secondName1 := getInitials("tifa lockhart")
	fmt.Println(firstName1, secondName1)

	firstName2, secondName2 := getInitials("cloud strife")
	fmt.Println(firstName2, secondName2)

	firstName3, secondName3 := getInitials("barret")
	fmt.Println(firstName3, secondName3)
}