package pratice

import "fmt"

func fmttest() {
	y := "asd"
	//Println 輸出一列
	fmt.Println(x)
	fmt.Println(x, x, x, x) //X是在1_var宣告的
	fmt.Println(x + x)
	fmt.Println(y, y, y)   //中間插空格
	fmt.Println(y + y + y) //相連
	//Printf 輸出格式化
	fmt.Printf("%d,%c,%s,%v\n", x, x, y, x) //%d: digit 10進位的數字) ;%c: char 字元 ;%s: string  (字串) ;%v: value   (值)
	fmt.Println(`\n`)                       //``可顯示跳脫字元
	//sprint 組合字串
	str1 := fmt.Sprintln(y, y, y) //Sprintln插空格
	fmt.Println(str1)
	str2 := fmt.Sprint(y, y, y) //Sprint連著
	fmt.Println(str2)
	//scanln 接字串
	scan1, scan2, scan3 := "", "", ""

	fmt.Scanln(&scan1, &scan2) //scanln 讀一排 若中間有空格 要多加變數
	fmt.Println(scan1, scan2)
	fmt.Scan(&scan3, &scan2) //差在哪?
	fmt.Println(scan3, scan2)
}
