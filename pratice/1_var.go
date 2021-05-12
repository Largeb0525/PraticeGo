package pratice

import "fmt"

func vartest() { //字首小寫=pravite 給同package用
	var a int //未宣告=0
	var b int = 16
	var c = 87
	d := 69
	fmt.Println(a, b, c, d)
	var i, j int = 8, 9
	e, f := 3, 5
	g, f := 8, 7 //會覆蓋
	fmt.Println(i, j, e, f, g)

	fmt.Println(&a) //記憶體位置 類似static

	//const GRADE = 87 //常數 無法變動 命名通常全大寫
}

var x = 123 // 全域變數 宣告全域變數只能用var，無法用:=

func pointtest() {
	fmt.Println(&x)

	var x = 123 // 在`func區塊`。全新。
	fmt.Println(&x)

	// var x, y = 123, 100 // 因為x已經宣告過了，不能再用`var`。但以下卻可以
	x, y := 123, 100 // 在`func區塊`。此時x是上面宣告過的區域變數，而y是新的
	fmt.Println(&x, &y)

	if true {
		var x, y = 123, 100 // 在新區域`if區塊`。此時x跟y都是全新的。
		// 可以換成短變數宣告。

		fmt.Println(&x, &y)
	}

	fmt.Println(&x) // 脫離if區域，回到`func區塊`
}
