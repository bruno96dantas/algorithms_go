package main

import "fmt"
import "algorithms_go/binary_search"

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binary_search.BinarySearch(items, 63))
}
