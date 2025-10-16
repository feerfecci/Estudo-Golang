package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []string{"a", "b", "c", "d"}
	newInts := reverse(slice1)
	newString := reverse(slice2)
	fmt.Println(newInts)
	fmt.Println(newString)
}

type ConstraintCustom interface {
	int | string
}

//[T int | string] usado para passar todos os tipo aceitaveis para esse método
func reverse[T ConstraintCustom](slice []T) []T {
	newInt := make([]T, len(slice))

	newIntLen := len(slice) - 1
	for i := 0; i < len(slice); i++ {
		newInt[newIntLen] = slice[i]
		newIntLen--
	}

	return newInt
}
