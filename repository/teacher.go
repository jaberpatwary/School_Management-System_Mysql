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
