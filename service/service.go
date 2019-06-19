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
