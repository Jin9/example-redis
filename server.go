package main

import (
	"wasabi/router"
)

func main() {
	e := router.Init()
	e.Logger.Fatal(e.Start(":8901"))
}
