package pratice

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LoginInfo struct {
	UserID     int64     `json:"userId"`
	ClientIP   string    `json:"clientIP"`
	LoginState string    `json:"loginState"`
	LoginTime  time.Time `json:"loginTime"`
}

func mongotest() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	db := session.DB("test")
	c := db.C("abc")

	err = c.Insert(NewLoginInfo(1234, "127.0.0.1", "success"))
	if err != nil {
		panic(err)
	}

	var loginHistory []LoginInfo
	err = c.Find(bson.M{"userid": 1234}).All(&loginHistory)
	if err != nil {
		panic(err)
	}
	for _, history := range loginHistory {
		fmt.Println(history)
	}

}

func NewLoginInfo(id int64, clientIP string, loginState string) *LoginInfo {
	return &LoginInfo{
		UserID:     id,
		ClientIP:   clientIP,
		LoginState: loginState,
		LoginTime:  time.Now(),
	}
}
