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
