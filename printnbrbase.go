package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	// The base has to be at least 2 characters
	if len(base) < 2 {
		fmt.Println("NV")
	}

	// We check if all the characters in base are unique, if they are not, we print "NV"
	stop := false
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				stop = true
			}
		}
	}
	if stop {
		fmt.Println("NV")
	} else {
		// If the number is negative, we multiply by -1 and later print the '-' sign
		isNegative := false
		if nbr < 0 {
			nbr *= -1
			isNegative = true
		}
		// Make nbr into an int slice by filling up an empty slice with the digits one by one
		nbrSlice := make([]int, 0)
		baseSlice := []rune(base)
		varTemp := nbr
		length := 0
		for varTemp > 0 {
			nbrSlice = append(nbrSlice, varTemp%len(base))
			varTemp /= len(base)
			length++
		}
		// Putting the integers into the rune slice in reverse order for printing
		sliceToPrint := make([]rune, length)
		for index := len(sliceToPrint) - 1; index >= 0; index-- {
			sliceToPrint[index] = baseSlice[nbrSlice[index]]
		}
		if isNegative {
			z01.PrintRune('-')
		}
		for i := len(sliceToPrint) - 1; i >= 0; i-- {
			z01.PrintRune(sliceToPrint[i])
		}
	}
}
