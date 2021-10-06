package main

import (
	"fmt"
	"strings"
	"sort"
)

func main()  {

	// strings
	var nameOne string = "meina"
	var nameTwo = "sarah"
	nameThree := "lol"
	fmt.Println(nameOne, nameTwo, nameThree)

	// integers
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint16 = 256
	fmt.Println(numOne, numTwo, numThree)

	// Print
	fmt.Print("Hello, ")
	fmt.Print("World!\n")

	// Println
	fmt.Println("My name is", nameOne, "and my age is", ageTwo)

	// Printf(formatted strings) %_ = format specifier
	// fmt.Printf("My name is %v and my age is %v", nameOne, ageTwo)

	// Sprintf(save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", ageTwo, nameOne)
	fmt.Println("the saved string is:", str)

	// Arrays
	var ages = [3]int{20, 25, 30}
	names := [4]string{"meina", "sarah", "yoyo", "Alex"}
	names[1] = "luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices(use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	// Standard Library
	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "th"))
	fmt.Println(strings.Split(greeting, " "))

	// the original value is unchanged
	fmt.Println("original string value = ", greeting)

	nums := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(nums)
	index := sort.SearchInts(nums, 30)
	fruits := []string{"apple", "peach", "avocado", "banana"}
	sort.Strings(fruits)

	fmt.Println(nums)
	fmt.Println(index)
	fmt.Println(fruits)
	fmt.Println(sort.SearchStrings(fruits, "apple"))


}