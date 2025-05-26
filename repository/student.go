package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramadhanalfarisi/go-simple-crud/model"
)

type StudentRepository struct {
	Db *sql.DB
}

func NewStudentRepository(db *sql.DB) StudentRepositoryInterface {
	return &StudentRepository{Db: db}
}

// Insert student
func (p *StudentRepository) InsertStudent(post model.PostStudent) bool {
	stmt, err := p.Db.Prepare("INSERT INTO student(name, date_of_birth, email) VALUES (?,?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Name, post.Date_Of_Birth, post.Email)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}
