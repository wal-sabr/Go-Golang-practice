package main

import "fmt"

func linearsearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	items := []int{62, 27, 420, 9, 100, 367}
	fmt.Println(linearsearch(items, 420))
}
