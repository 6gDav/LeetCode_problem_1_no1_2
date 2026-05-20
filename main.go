package main

import "fmt"

func main() {
	list := []int{3, 2, 4}
	var target int

	fmt.Print("Gimme a target number: ")
	fmt.Scanf("%v", &target)

	res := twoSum(list, target)
	fmt.Println(res)
}
