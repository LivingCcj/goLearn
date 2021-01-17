package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{2,31,41,15,15,2,32,0}
	sort.Ints(a)

	for i,v := range a{
		fmt.Println(i ,v)
	}
}
