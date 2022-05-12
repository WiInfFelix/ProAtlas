package exercises

import (
	database "github.com/WiInfFelix/proatlas/internal/pkg/db/mysql"
	"github.com/WiInfFelix/proatlas/internal/users"
	"log"
)

type Exercise struct {
	ID           string
	Name         string
	Description  string
	TargetMuscle string
	Creator      *users.User
}

func (exec Exercise) Save() int64 {

	stmt, err := database.Db.Prepare("INSERT INTO Exercises(ExerciseName, ExerciseDescription, TargetMuscle) VALUES (?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(exec.Name, exec.Description, exec.TargetMuscle)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func GetAll() []Exercise {

	stmt, err := database.Db.Prepare("Select id, ExerciseName, ExerciseDescription, TargetMuscle from Exercises")
	if err != nil {
		log.Panic(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Panic(err)
	}

	var exercises []Exercise

	for rows.Next() {
		var exercise Exercise
		err := rows.Scan(&exercise.ID, &exercise.Name, &exercise.Description, &exercise.TargetMuscle)
		if err != nil {
			log.Fatal(err)
		}
		exercises = append(exercises, exercise)

	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return exercises

}
