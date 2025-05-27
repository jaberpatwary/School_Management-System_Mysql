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

// GetAll student
func (p *StudentRepository) GetAllStudent() []model.Student {
	query, err := p.Db.Query("SELECT * FROM school.student")
	if err != nil {
		log.Println(err)
		return nil
	}
	var students []model.Student
	if query != nil {
		for query.Next() {
			var (
				id            int
				name          string
				date_of_birth *int
				email         *string
			)
			err := query.Scan(&id, &name, &date_of_birth, &email)
			if err != nil {
				log.Println(err)
			}
			student := model.Student{ID: id, Name: name, Date_Of_Birth: date_of_birth, Email: email}
			students = append(students, student)
		}
	}
	return students
}

// GetOne student
func (m *StudentRepository) GetOneStudent(id int) model.Student {
	query, err := m.Db.Query("SELECT * FROM school.student WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return model.Student{}
	}
	defer query.Close()
	var student model.Student
	if query != nil {
		for query.Next() {
			var (
				id            int
				name          string
				date_of_birth *int
				email         *string
			)
			err := query.Scan(&id, &name, &date_of_birth, &email)
			if err != nil {
				log.Println(err)
				return model.Student{}
			}
			student = model.Student{ID: id, Name: name, Date_Of_Birth: date_of_birth, Email: email}
		}
	}
	return student
}

//update student

func (p *StudentRepository) UpdateStudent(id int, post model.PostStudent) model.Student {
	_, err := p.Db.Exec("UPDATE student SET name = ?, date_of_birth = ?, email = ? WHERE id = ?", post.Name, post.Date_Of_Birth, post.Email, id)
	if err != nil {
		log.Println(err)
		return model.Student{}
	}
	return p.GetOneStudent(id)
}
