package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramadhanalfarisi/go-simple-crud/model"
)

type CourseRepository struct {
	Db *sql.DB
}

func NewCourseRepository(db *sql.DB) CourseRepositoryInterface {
	return &CourseRepository{Db: db}
}

// Insert course
func (p *CourseRepository) InsertCourse(post model.PostCourse) bool {
	stmt, err := p.Db.Prepare("INSERT INTO course(course_name, department, credit) VALUES (?,?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Course_Name, post.Department, post.Credit)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}
