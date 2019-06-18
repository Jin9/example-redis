package main

import (
	"wasabi/router"
)

func main() {
	e := router.Init()
	e.Start(":8901")
}
