package database

import (
	"backend/models"
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type StudentInfo = models.StudentInfo

var db *sql.DB

func OpenDB(path string) error {
	var err error
	db, err = sql.Open("sqlite", path)
	if err != nil {
		return err
	}
	return nil
}

func RunMigrations() error {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS turmas (
		turma_id TEXT PRIMARY KEY,
		name TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS students (
		student_id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		matricula TEXT UNIQUE NOT NULL,
		turma_id TEXT NOT NULL,
		FOREIGN KEY (turma_id)
			REFERENCES turmas (turma_id)
	);
	`

	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migrations completed")
	return nil
}

func WriteStudentInfo(students []StudentInfo, turma_id string) error {

	stmt, err := db.Prepare("INSERT INTO students(student_id, name, matricula, turma_id) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, student := range students {
		_, err = stmt.Exec(student.Id, student.Name, student.Matricula, turma_id)
		if err != nil {
			return err
		}
	}
	return nil
}

func FetchStudentinfo(classId string) ([]StudentInfo, error) {
	rows, err := db.Query("SELECT student_id, name, matricula FROM students WHERE turma_id = ?", classId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []StudentInfo
	for rows.Next() {
		var student StudentInfo
		if err := rows.Scan(&student.Id, &student.Name, &student.Matricula); err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func CleanUp(classId string) error {
	query := `DELETE FROM students
				WHERE turma_id = ?;`

	_, err := db.Exec(query, classId)
	return err
}
