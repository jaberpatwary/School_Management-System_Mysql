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

// GetAll enroll
func (p *EnrollRepository) GetAllEnroll() []model.Enroll {
	query, err := p.Db.Query("SELECT * FROM school.enroll")
	if err != nil {
		log.Println(err)
		return nil
	}
	var enrolls []model.Enroll
	if query != nil {
		for query.Next() {
			var (
				id         int
				course_id  int
				student_id *int
				date       *int
			)
			err := query.Scan(&id, &course_id, &student_id, &date)
			if err != nil {
				log.Println(err)
			}
			enroll := model.Enroll{ID: id, Course_Id: course_id, Student_Id: student_id, Date: date}
			enrolls = append(enrolls, enroll)
		}
	}
	return enrolls
}
