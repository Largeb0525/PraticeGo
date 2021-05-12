package pratice

import "fmt"

type money struct {
	Blue1000 int // 藍色小朋友
	Red100   int // 紅色國父
	Card     string
}

type bag struct {
	money  // 直接放入結構就好
	iphone bool
	key    string
}

func maptest() {
	var sex = map[string]string{
		"male":    "男",
		"female":  "女",
		"shemale": "中性",
	}
	var yesno = map[bool]int{
		true:  1,
		false: 0,
	}
	fmt.Println(sex["male"], yesno[true])
	for key, value := range sex {
		fmt.Println(key, value)
	}
}

func structtest() {

	var bag1 = bag{
		money: money{
			Card:   "visa",
			Red100: 6,
		},
		iphone: true,
		key:    "yes",
	}

	fmt.Println(bag1)

}
