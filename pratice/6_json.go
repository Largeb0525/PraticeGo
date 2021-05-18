package pratice

import (
	"fmt"
	"log"
	"encoding/json"
)

type wow struct{
	ID  string `json:"Id"` //``內可以更改Key輸出
	Job string
	Level int
	Troll bool
	test string  //Key字首要大寫 否則無法Marshal 如test無法輸出
}

func jsontest()  {
	player1:=wow{
		ID:"Largeb",
		Job:"Duird",
		Level:60,
		Troll:false,
		test:"None",
	}

	out, err := json.Marshal(player1)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string(out))
	
	playercopy:=wow{}
	
	err = json.Unmarshal(out, &playercopy)
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("u: %+v\n", playercopy)
}