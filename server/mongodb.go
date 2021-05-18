package server

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LoginInfo struct {
	UserID    string `json:"userId"`
	Action    string
	Money     int
	Total     int
	LoginTime time.Time
}

var db *mgo.Database
var coll *mgo.Collection

func mongointi() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	db = session.DB("test")
	coll = db.C("abc")

}

func newdata(id string, act string, money int) *LoginInfo {
	return &LoginInfo{
		UserID:    id,
		Action:    act,
		Money:     money,
		LoginTime: time.Now(),
	}
}

func deposit(id string, act string, money int) {
	coll.Insert(newdata(id, act, money))
}

func withdraw(id string, act string, money int) {
	coll.Insert(newdata(id, act, money))
}

func balance(id string) int {
	var loginHistory []LoginInfo
	err3 := coll.Find(bson.M{"userid": id}).All(&loginHistory)
	if err3 != nil {
		panic(err3)
	}
	sum := 0
	for _, history := range loginHistory {
		if history.Action == "存錢" {
			sum += history.Money
		} else if history.Action == "領錢" {
			sum -= history.Money
		}
	}
	return sum
}
