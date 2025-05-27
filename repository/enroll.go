package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramadhanalfarisi/go-simple-crud/model"
)

type EnrollRepository struct {
	Db *sql.DB
}

func NewEnrollRepository(db *sql.DB) EnrollRepositoryInterface {
	return &EnrollRepository{Db: db}
}

// Insert enroll
func (p *EnrollRepository) InsertEnroll(post model.PostEnroll) bool {
	stmt, err := p.Db.Prepare("INSERT INTO enroll(course_id, student_id, date) VALUES (?,?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Course_Id, post.Student_Id, post.Date)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}
