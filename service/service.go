package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"wasabi/constant"
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

func hashValue(keyValue string) (string, error) {
	h := sha3.New256()

	_, err := h.Write([]byte(keyValue))

	if err != nil {
		log.Println(err.Error())
		return "", errors.New("Cannot hash value")
	}

	keyBytes := h.Sum(nil)
	return fmt.Sprintf("%x", keyBytes), nil
}

func prepareUToken(username string) (string, error) {
	keyValue := fmt.Sprintf("%v%v", generateToken(), username)
	hashKeyValue, err := hashValue(keyValue)
	if err != nil {
		return "", err
	}
	return hashKeyValue, nil
}

func processSaveUser(username string, ptoken string) error {
	utoken, err := prepareUToken(username)
	if err != nil {
		return err
	}
	if err = saveUser(username, utoken, ptoken); err != nil {
		return err
	}
	return nil
}

func saveUser(username string, utoken string, ptoken string) error {
	u := model.NewUser(utoken, ptoken, false)
	val, _ := json.Marshal(u)
	if err := db.SetData(username, val, 0, constant.UserDB); err != nil {
		return err
	}
	return nil
}

func validateUser(username string) error {
	val, err := db.GetExistsKey(username, constant.UserDB)
	if err != nil {
		return err
	}
	if val {
		log.Println("Username is alerady exists")
		return errors.New("Username is alerady exists")
	}
	return nil
}

// RegisterUser is a service for record new member
func RegisterUser(user *model.RegisterUserRequest) error {
	if err := validateUser(user.UserName); err != nil {
		return err
	}
	if err := db.RegisterNewUser(user); err != nil {
		return err
	}
	if err := processSaveUser(user.UserName, user.PToken); err != nil {
		return err
	}
	return nil
}

func matchUser(username string, ptoken string) (string, error) {
	val, err := db.GetData(username, constant.UserDB)
	if err != nil {
		log.Println(err)
		return "", errors.New("Invalid User or Password")
	}

	data := &model.User{}
	json.Unmarshal([]byte(val), &data)
	if data.PToken != ptoken {
		log.Println("password not match")
		return "", errors.New("Invalid User or Password")
	}
	if data.IsLogedIn {
		log.Println("this user is loged in")
		return "", errors.New("User is already loged in")
	}

	return data.UToken, nil
}

func createAccessToken(utoken string) (string, error) {
	atoken, err := hashValue(generateToken())
	if err != nil {
		return "", err
	}
	if err := db.SetData(atoken, utoken, 5*60*time.Second, constant.AccessDB); err != nil {
		return "", err
	}
	return atoken, nil
}

// LoginUser is a service for using login
func LoginUser(user *model.LoginUserRequest) (string, error) {
	utoken, err := matchUser(user.Username, user.PToken)
	if err != nil {
		return "", err
	}
	if err = db.SetData(utoken, user.Username, 0, constant.UserTokenDB); err != nil {
		return "", err
	}
	atoken, err := createAccessToken(utoken)
	if err != nil {
		return "", err
	}
	return atoken, nil
}

// LogOutUser is a service for using logout
func LogOutUser(atoken string) error {
	result, err := db.GetData(atoken, constant.AccessDB)
	if err != nil {
		if err.Error() != "Key is not exists" {
			return err
		}
	}
	fmt.Println(result)
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
