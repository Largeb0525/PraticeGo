package pratice

import "fmt"

func arraytest() {
	var arr1 [5]int = [5]int{1, 2, 3}
	arr2 := [3]int{4, 5, 6}
	arr3 := [...]string{ //...代表自動計算 分行句尾要,
		"asd",
		"qwe",
		"zxc",
	}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	slice1 := make([]int, 5, 10) //Variable := make([]Type, Len, Cap)
	fmt.Println(slice1)
	slice2 := []int{1, 2, 3} //[]內無數字代表slice
	slice2 = append(slice2, 4, 5, 6)
	fmt.Println(slice2)
}
