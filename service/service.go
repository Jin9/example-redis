package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"wasabi/db"
	"wasabi/model"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/sha3"
)

const (
	location = "Asia/Bangkok"
)

func getTimeStamp() int64 {
	timeLocation, _ := time.LoadLocation(location)
	now := time.Now().In(timeLocation)
	return now.Unix()
}

func generateToken() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}

func hashUUID(keyValue string) (string, error) {
	h := sha3.New256()

	_, err := h.Write([]byte(keyValue))

	if err != nil {
		return "", err
	}

	keyBytes := h.Sum(nil)
	return fmt.Sprintf("%x", keyBytes), nil
}

func saveUser(user *model.RegisterUserRequest) error {
	u := model.NewUser("", 0, user.Email, user.Phone)
	val, _ := json.Marshal(u)
	if err := db.SetData(user.UserName, val, 0); err != nil {
		return err
	}
	return nil
}

// RegisterUser is a service for record new member
func RegisterUser(user *model.RegisterUserRequest) error {
	val, err := db.GetExistsKey(user.UserName)
	if err != nil {
		return err
	}
	if val {
		return errors.New("This username is alerady exists")
	}
	// if err := db.RegisterNewUser(user); err != nil {
	// 	return err
	// }
	if err := saveUser(user); err != nil {
		return err
	}
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
