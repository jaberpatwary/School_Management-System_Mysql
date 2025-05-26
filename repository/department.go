package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramadhanalfarisi/go-simple-crud/model"
)

type DepartmentRepository struct {
	Db *sql.DB
}

func NewDepartmentRepository(db *sql.DB) DepartmentRepositoryInterface {
	return &DepartmentRepository{Db: db}
}

// Insert department
func (p *DepartmentRepository) InsertDepartment(post model.PostDepartment) bool {
	stmt, err := p.Db.Prepare("INSERT INTO department(name, Department_head) VALUES (?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Name, post.Department_Head)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// GetAll department
func (p *DepartmentRepository) GetAllDepartment() []model.Department {
	query, err := p.Db.Query("SELECT * FROM school.department")
	if err != nil {
		log.Println(err)
		return nil
	}
	var departments []model.Department
	if query != nil {
		for query.Next() {
			var (
				id              int
				name            string
				department_head *string
			)
			err := query.Scan(&id, &name, &department_head)
			if err != nil {
				log.Println(err)
			}
			department := model.Department{ID: id, Name: name, Department_Head: department_head}
			departments = append(departments, department)
		}
	}
	return departments
}

// GetOne department
func (m *DepartmentRepository) GetOneDepartment(id int) model.Department {
	query, err := m.Db.Query("SELECT * FROM school.department WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return model.Department{}
	}
	defer query.Close()
	var department model.Department
	if query != nil {
		for query.Next() {
			var (
				id              int
				name            string
				department_head *string
			)
			err := query.Scan(&id, &name, &department_head)
			if err != nil {
				log.Println(err)
				return model.Department{}
			}
			department = model.Department{ID: id, Name: name, Department_Head: department_head}
		}
	}
	return department
}

//update department

func (p *DepartmentRepository) UpdateDepartment(id int, post model.PostDepartment) model.Department {
	_, err := p.Db.Exec("UPDATE department SET name = ?, department_head = ? WHERE id = ?", post.Name, post.Department_Head, id)
	if err != nil {
		log.Println(err)
		return model.Department{}
	}
	return p.GetOneDepartment(id)
}
