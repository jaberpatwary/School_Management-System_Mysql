package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramadhanalfarisi/go-simple-crud/model"
)

type teacherRepository struct {
	Db *sql.DB
}

func NewTeacherRepository(db *sql.DB) TeacherRepositoryInterface {
	return &teacherRepository{Db: db}
}

// Insert teacher
func (p *teacherRepository) InsertTeacher(post model.PostTeacher) bool {
	stmt, err := p.Db.Prepare("INSERT INTO teacher(name, Department_id, email) VALUES (?,?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Name, post.Department_Id, post.Email)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// GetAll teacher
func (p *teacherRepository) GetAllTeacher() []model.Teacher {
	query, err := p.Db.Query("SELECT * FROM school.teacher")
	if err != nil {
		log.Println(err)
		return nil
	}
	var teachers []model.Teacher
	if query != nil {
		for query.Next() {
			var (
				id            int
				name          string
				department_id *int
				email         *string
			)
			err := query.Scan(&id, &name, &department_id, &email)
			if err != nil {
				log.Println(err)
			}
			teacher := model.Teacher{ID: id, Name: name, Department_Id: department_id, Email: email}
			teachers = append(teachers, teacher)
		}
	}
	return teachers
}

// GetOne teacher
func (m *teacherRepository) GetOneTeacher(id int) model.Teacher {
	query, err := m.Db.Query("SELECT * FROM school.teacher WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return model.Teacher{}
	}
	defer query.Close()
	var teacher model.Teacher
	if query != nil {
		for query.Next() {
			var (
				id            int
				name          string
				department_id *int
				email         *string
			)
			err := query.Scan(&id, &name, &department_id, &email)
			if err != nil {
				log.Println(err)
				return model.Teacher{}
			}
			teacher = model.Teacher{ID: id, Name: name, Department_Id: department_id, Email: email}
		}
	}
	return teacher
}

//update teacher

func (p *teacherRepository) UpdateTeacher(id int, post model.PostTeacher) model.Teacher {
	_, err := p.Db.Exec("UPDATE teacher SET name = ?, department_id = ?, email = ? WHERE id = ?", post.Name, post.Department_Id, post.Email, id)
	if err != nil {
		log.Println(err)
		return model.Teacher{}
	}
	return p.GetOneTeacher(id)
}

// Delete teacher
func (m *teacherRepository) DeleteTeacher(id int) bool {
	_, err := m.Db.Exec("DELETE FROM teacher WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
