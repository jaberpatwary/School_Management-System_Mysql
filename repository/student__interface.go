package repository

import "github.com/ramadhanalfarisi/go-simple-crud/model"

type StudentRepositoryInterface interface {
	InsertStudent(model.PostStudent) bool
	GetAllStudent() []model.Student
	GetOneStudent(int) model.Student
	//DeleteStudent(int) bool
	//UpdateStudent(int, model.PostStudent) model.Student
}
