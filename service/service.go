package service

import (
	"fmt"
	"time"
)

const (
	location = "Asia/Bangkok"
)

var (
	timeLocation *time.Location
)

// InitTimeLocation is used for get local time
func InitTimeLocation() {
	tl, _ := time.LoadLocation(location)
	timeLocation = tl

	timeNow()
}

func timeNow() {
	now := time.Now().In(timeLocation)
	sec := now.Unix()
	nano := now.UnixNano()
	fmt.Println(now)
	fmt.Println(sec)
	fmt.Println(nano)
}

func mashal() {
	// u := &User{
	// 	FirstName: "Adam",
	// 	LastName:  "Smith",
	// }

	// v, _ := json.Marshal(u)

	// db.SetData("key", v, 3*time.Second)
	// val := db.GetData("key")

	// data := &User{}
	// json.Unmarshal([]byte(val), &data)

	// fmt.Println(data)
}
