package repository

import "github.com/ramadhanalfarisi/go-simple-crud/model"

type CourseRepositoryInterface interface {
	InsertCourse(model.PostCourse) bool
	GetAllCourse() []model.Course
	GetOneCourse(int) model.Course
	//DeleteCourse(int) bool
	//UpdateCourse(int, model.PostCourse) model.Course
}
