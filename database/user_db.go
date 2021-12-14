package database

import (
	"ProAtlasV6/models"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	"log"
)

type DBError struct {
	ErrString string
}

func InsertUserIntoDB(user models.User) (err error) {

	insertStatement, err := GetDB().Prepare("INSERT INTO users (Username, Email, Password) VALUES (?,?,?)")

	res, err := insertStatement.Exec(user.Username, user.Email, user.Password)

	if err != nil {

		if sqlLiteErr, ok := err.(sqlite3.Error); ok {
			if sqlLiteErr.Code == sqlite3.ErrConstraint {
				err = errors.New("there was an error concerning unique constraints")
			}
		}
		log.Printf("Error inserting into DB... \n %s \n", err)
		return err
	}

	insertID, err := res.LastInsertId()

	log.Printf("Inserted entry with %s", insertID)

	return
}

func GetAllUsers() (err error, users []models.User) {

	rows, err := GetDB().Query("SELECT users.ID, users.Username, users.Email FROM users")
	if err != nil {
		log.Printf("There was an error querying all users: %s \n", err)
		return err, users
	}

	for rows.Next() {
		var userDummy models.User
		err = rows.Scan(&userDummy.Id, &userDummy.Username, &userDummy.Email)
		if err != nil {
			log.Println(err)
		}
		users = append(users, userDummy)
	}

	return
}

func GetUserByID(userID int) (user models.User, err error) {

	query := "SELECT users.ID, users.Username, users.Email FROM users WHERE id=$1"

	row := GetDB().QueryRow(query, userID)

	err = row.Scan(&user.Id, &user.Username, &user.Email)
	if err != nil {
		fmt.Println("No rows found for ID")
		return
	}

	return
}
