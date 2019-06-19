package main

import (
	// "wasabi/service"
	// "wasabi/router"
	"encoding/json"
	"fmt"
	"time"
	"wasabi/db"
)

// User is
type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// UnmarshalBinary is
func (m *User) UnmarshalBinary(data []byte) error {
	// convert data to yours, let's assume its json data
	return json.Unmarshal(data, m)
}

func main() {
	// e := router.Init()
	// e.Logger.Fatal(e.Start(":8880"))

	u := &User{
		FirstName: "Adam",
		LastName:  "Smith",
	}

	v, _ := json.Marshal(u)

	db.SetData("key", v, 3*time.Second)
	val := db.GetData("key")

	data := &User{}
	json.Unmarshal([]byte(val), &data)

	fmt.Println(data)
}
