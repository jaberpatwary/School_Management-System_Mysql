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

// GetAll course
func (p *CourseRepository) GetAllCourse() []model.Course {
	query, err := p.Db.Query("SELECT * FROM school.course")
	if err != nil {
		log.Println(err)
		return nil
	}
	var courses []model.Course
	if query != nil {
		for query.Next() {
			var (
				id          int
				course_name string
				department  *string
				credit      *string
			)
			err := query.Scan(&id, &course_name, &department, &credit)
			if err != nil {
				log.Println(err)
			}
			course := model.Course{ID: id, Course_Name: course_name, Department: department, Credit: credit}
			courses = append(courses, course)
		}
	}
	return courses
}

// GetOne course
func (m *CourseRepository) GetOneCourse(id int) model.Course {
	query, err := m.Db.Query("SELECT * FROM school.course WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return model.Course{}
	}
	defer query.Close()
	var course model.Course
	if query != nil {
		for query.Next() {
			var (
				id          int
				course_name string
				department  *string
				credit      *string
			)
			err := query.Scan(&id, &course_name, &department, &credit)
			if err != nil {
				log.Println(err)
				return model.Course{}
			}
			course = model.Course{ID: id, Course_Name: course_name, Department: department, Credit: credit}
		}
	}
	return course
}

//update course

func (p *CourseRepository) UpdateCourse(id int, post model.PostCourse) model.Course {
	_, err := p.Db.Exec("UPDATE course SET course_name = ?, department = ?, credit = ? WHERE id = ?", post.Course_Name, post.Department, post.Credit, id)
	if err != nil {
		log.Println(err)
		return model.Course{}
	}
	return p.GetOneCourse(id)
}
