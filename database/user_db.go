package database

import (
	"ProAtlasV6/models"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type DBError struct {
	ErrString string
}

func InsertUserIntoDB(user models.User) (err error) {

	insertStatement, err := GetDB().Prepare("INSERT INTO users (Username, Email, Password) VALUES (?,?,?)")

	res, err := insertStatement.Exec(user.Username, user.Email, user.Password)

	if err != nil {

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

	query, err := GetDB().Prepare("SELECT users.ID, users.Username, users.Email FROM users WHERE id=?")
	log.Printf("Querying user for id %d", userID)

	err = query.QueryRow(userID).Scan(&user.Id, &user.Username, &user.Email)
	if err != nil {
		fmt.Printf("No rows found for ID: %s \n")
		return
	}

	return
}

func DeleteUser(userID int) (err error) {
	query, err := GetDB().Prepare("DELETE FROM users WHERE Id =?")
	if err != nil {
		return err
	}

	_, err = query.Exec(userID)
	if err != nil {
		return err
	}

	return
}
