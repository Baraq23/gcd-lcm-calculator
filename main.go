package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <num1> <num2>")
		return
	}
	num1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: Enter only numbers")
	}
	num2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: Enter only numbers")
	}

	fmt.Printf("Num1 & Num2: %v %v\n", num1, num2)

	gcd, lcm := GcdLcm(num1, num2)
	fmt.Printf("The GCD is: %3d\n", gcd)
	fmt.Printf("The LCM is: %3d\n", lcm)
}

// Function GcdLcm() takes two numbers and returns their lcm and gcd
func GcdLcm(a, b int) (gcd, lcm int) {
	var lg, sm int

	if a > b {
		lg = a
		sm = b
	} else {
		lg = b
		sm = a
	}

	// for lcm logically
	l := lg
	for i := 2; i != 1; i++ {
		if l%sm == 0 {
			lcm = l
			break
		}
		l += lg
	}

	// fmt.Println("large & small ", lg, sm)

	// for GCD
	for i := sm; i != 0; i-- {
		if sm%i == 0 && lg%i == 0 {
			gcd = i
			break
		}
	}

	// for lcm formular
	// lcm = (a*b)/gcd
	return
}
