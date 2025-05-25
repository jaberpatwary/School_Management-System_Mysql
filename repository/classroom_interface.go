package repository

import "github.com/ramadhanalfarisi/go-simple-crud/model"

type ClassroomRepositoryInterface interface {
	InsertClassroom(model.PostClassroom) bool
	GetAllClassroom() []model.Classroom
	GetOneClassroom(int) model.Classroom
	//DeleteClassroom(int) bool
	//UpdateClassroom(int, model.PostClassroom) model.Classroom
}
