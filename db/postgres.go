package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"wasabi/model"

	// pq is a postgres driver
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	pwd    = "password"
	dbname = "profile"
)

func connectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pwd, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// RegisterNewUser is used for insert new user into db
func RegisterNewUser(user *model.RegisterUserRequest) error {
	db, err := connectDB()
	if err != nil {
		log.Println(err.Error())
		return errors.New("Cannat connect to database")
	}

	defer db.Close()
	insertStmt := `INSERT INTO users.members 
						(p_token, username, firstname, lastname, email, phone) 
				   VALUES 
						($1, $2, $3, $4, $5, $6);`

	_, err = db.Exec(insertStmt, user.PToken, user.UserName, user.FirstName, user.LastName, user.Email, user.Phone)
	if err != nil {
		log.Println(err.Error())
		return errors.New("Cannot register new user")
	}

	return nil
}
