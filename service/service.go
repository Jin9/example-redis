package service

import (
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
}
