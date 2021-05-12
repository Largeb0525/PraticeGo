package pratice

import "fmt"

func conditiontest() {
	a, b, c := 1, 2, 3
	if a > 0 {
		fmt.Println("87")
	} else if a < 0 {
		fmt.Println(b)
	} else {
		fmt.Println(c)
	}

	for i := 0; i < 3; i++ {
		fmt.Print(i)
	}
	fmt.Println("")

	for a < 3 { //相當於while
		a++
		fmt.Print(a)
	}
	fmt.Println("")

	intarr := []int{4, 5, 6, 7}
	for index, unit := range intarr { //Range 掃全array或map
		fmt.Println(index, unit)
	}
}
