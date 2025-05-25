package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramadhanalfarisi/go-simple-crud/model"
)

type ClassroomRepository struct {
	Db *sql.DB
}

func NewClassroomRepository(db *sql.DB) ClassroomRepositoryInterface {
	return &ClassroomRepository{Db: db}
}

/* Insert category
func (p *CategoryRepository) InsertCategory(post model.PostCategory) bool {
	stmt, err := p.Db.Prepare("INSERT INTO category(name, types) VALUES (?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Name, post.Types)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// GetAll category
func (p *CategoryRepository) GetAllCategory() []model.Category {
	query, err := p.Db.Query("SELECT * FROM eco.category")
	if err != nil {
		log.Println(err)
		return nil
	}
	var categorys []model.Category
	if query != nil {
		for query.Next() {
			var (
				id    int
				name  string
				types *string
			)
			err := query.Scan(&id, &name, &types)
			if err != nil {
				log.Println(err)
			}
			category := model.Category{ID: id, Name: name, Types: types}
			categorys = append(categorys, category)
		}
	}
	return categorys
}

// GetOne category
func (m *CategoryRepository) GetOneCategory(id int) model.Category {
	query, err := m.Db.Query("SELECT * FROM eco.category WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return model.Category{}
	}
	defer query.Close()
	var category model.Category
	if query != nil {
		for query.Next() {
			var (
				id    int
				name  string
				types *string
			)
			err := query.Scan(&id, &name, &types)
			if err != nil {
				log.Println(err)
				return model.Category{}
			}
			category = model.Category{ID: id, Name: name, Types: types}
		}
	}
	return category
}

//update category

func (p *CategoryRepository) UpdateCategory(id int, post model.PostCategory) model.Category {
	_, err := p.Db.Exec("UPDATE category SET name = ?, types = ? WHERE id = ?", post.Name, post.Types, id)
	if err != nil {
		log.Println(err)
		return model.Category{}
	}
	return p.GetOneCategory(id)
}

// Delete category
func (m *CategoryRepository) DeleteCategory(id int) bool {
	_, err := m.Db.Exec("DELETE FROM category WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}*/

// Insert classroom
func (p *ClassroomRepository) InsertClassroom(post model.PostClassroom) bool {
	stmt, err := p.Db.Prepare("INSERT INTO classroom(room, capacity) VALUES (?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Room, post.Capacity)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// GetAll classroom
func (p *ClassroomRepository) GetAllClassroom() []model.Classroom {
	query, err := p.Db.Query("SELECT * FROM school.classroom")
	if err != nil {
		log.Println(err)
		return nil
	}
	var classrooms []model.Classroom
	if query != nil {
		for query.Next() {
			var (
				id       int
				room     int
				capacity *int
			)
			err := query.Scan(&id, &room, &capacity)
			if err != nil {
				log.Println(err)
			}
			classroom := model.Classroom{ID: id, Room: room, Capacity: capacity}
			classrooms = append(classrooms, classroom)
		}
	}
	return classrooms
}

// GetOne classroom
func (m *ClassroomRepository) GetOneClassroom(id int) model.Classroom {
	query, err := m.Db.Query("SELECT * FROM school.classroom WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return model.Classroom{}
	}
	defer query.Close()
	var classroom model.Classroom
	if query != nil {
		for query.Next() {
			var (
				id       int
				room     int
				capacity *int
			)
			err := query.Scan(&id, &room, &capacity)
			if err != nil {
				log.Println(err)
				return model.Classroom{}
			}
			classroom = model.Classroom{ID: id, Room: room, Capacity: capacity}
		}
	}
	return classroom
}

//update classroom

func (p *ClassroomRepository) UpdateClassroom(id int, post model.PostClassroom) model.Classroom {
	_, err := p.Db.Exec("UPDATE classroom SET room = ?, capacity = ? WHERE id = ?", post.Room, post.Capacity, id)
	if err != nil {
		log.Println(err)
		return model.Classroom{}
	}
	return p.GetOneClassroom(id)
}
