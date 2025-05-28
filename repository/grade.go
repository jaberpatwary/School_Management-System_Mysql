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

// GetAll grade
func (p *GradeRepository) GetAllGrade() []model.Grade {
	query, err := p.Db.Query("SELECT * FROM school.grade")
	if err != nil {
		log.Println(err)
		return nil
	}
	var grades []model.Grade
	if query != nil {
		for query.Next() {
			var (
				id         int
				course_id  int
				student_id *int
				score      *int
			)
			err := query.Scan(&id, &course_id, &student_id, &score)
			if err != nil {
				log.Println(err)
			}
			grade := model.Grade{ID: id, Course_Id: course_id, Student_Id: student_id, Score: score}
			grades = append(grades, grade)
		}
	}
	return grades
}

// GetOne grade
func (m *GradeRepository) GetOneGrade(id int) model.Grade {
	query, err := m.Db.Query("SELECT * FROM school.grade WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return model.Grade{}
	}
	defer query.Close()
	var grade model.Grade
	if query != nil {
		for query.Next() {
			var (
				id         int
				course_id  int
				student_id *int
				score      *int
			)
			err := query.Scan(&id, &course_id, &student_id, &score)
			if err != nil {
				log.Println(err)
				return model.Grade{}
			}
			grade = model.Grade{ID: id, Course_Id: course_id, Student_Id: student_id, Score: score}
		}
	}
	return grade
}

//update grade

func (p *GradeRepository) UpdateGrade(id int, post model.PostGrade) model.Grade {
	_, err := p.Db.Exec("UPDATE grade SET course_id = ?, student_id = ?, score = ? WHERE id = ?", post.Course_Id, post.Student_Id, post.Score, id)
	if err != nil {
		log.Println(err)
		return model.Grade{}
	}
	return p.GetOneGrade(id)
}
