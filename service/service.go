package service

import (
	"time"
)

const (
	location = "Asia/Bangkok"
)

func getTimeStamp() int64 {
	timeLocation, _ := time.LoadLocation(location)
	now := time.Now().In(timeLocation)
	return now.Unix()
}

// RegisterUser is a service for record new member
func RegisterUser() error {
	return nil
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
