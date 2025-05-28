package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramadhanalfarisi/go-simple-crud/model"
)

type GradeRepository struct {
	Db *sql.DB
}

func NewGradeRepository(db *sql.DB) GradeRepositoryInterface {
	return &GradeRepository{Db: db}
}

// Insert enroll
func (p *GradeRepository) InsertGrade(post model.PostGrade) bool {
	stmt, err := p.Db.Prepare("INSERT INTO grade(course_id, student_id, score) VALUES (?,?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Course_Id, post.Student_Id, post.Score)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}
