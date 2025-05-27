package repository

import "github.com/ramadhanalfarisi/go-simple-crud/model"

type EnrollRepositoryInterface interface {
	InsertEnroll(model.PostEnroll) bool
	//GetAllEnroll() []model.Enroll
	//GetOneEnroll(int) model.Enroll
	//DeleteEnroll(int) bool
	//UpdateEnroll(int, model.PostEnroll) model.Enroll
}
