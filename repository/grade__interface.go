package repository

import "github.com/ramadhanalfarisi/go-simple-crud/model"

type GradeRepositoryInterface interface {
	InsertGrade(model.PostGrade) bool
	GetAllGrade() []model.Grade
	GetOneGrade(int) model.Grade
	//DeleteGrade(int) bool
	//UpdateGrade(int, model.PostGrade) model.Grade
}
