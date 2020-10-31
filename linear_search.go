// Linear Search in Golang
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
	items := []int{95, 78, 46, 58, 4, 86, 99, 25, 30, 0, 9, 98776}
	fmt.Println(linearsearch(items, 9))
}
