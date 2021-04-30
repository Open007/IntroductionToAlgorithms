package main

import (
	"fmt"
)

func main() {
	var array [20]int

	for ii := range array {
		array[ii] = ii
		fmt.Print(array[ii])
	}
	fmt.Print("\n")

	sliceA := array[0:10]
	sliceB := append(sliceA, 0)

	for ii := range sliceA {
		fmt.Print(sliceA[ii])
	}
	fmt.Print("\n")

	for ii := range sliceB {
		fmt.Print(sliceB[ii])
	}
	fmt.Print("\n")

	for ii := range array {
		array[ii] = ii
		fmt.Print(array[ii])
	}
	fmt.Print("\n")

}
