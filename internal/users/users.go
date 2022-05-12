package users

import database "github.com/WiInfFelix/proatlas/internal/pkg/db/mysql"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user User) Create() {
	statement, err := database.Db.Prepare("INSERT INTO Users(Username, Password) VALUES (?,?)")

	print()
}
