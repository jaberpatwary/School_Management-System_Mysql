package repository

import "github.com/ramadhanalfarisi/go-simple-crud/model"

type TeacherRepositoryInterface interface {
	InsertTeacher(model.PostTeacher) bool
	GetAllTeacher() []model.Teacher
	GetOneTeacher(int) model.Teacher
	//DeleteTeacher(int) bool
	//UpdateTeacher(int, model.PostTeacher) model.Teacher
}
