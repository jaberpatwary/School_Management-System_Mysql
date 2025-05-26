package repository

import "github.com/ramadhanalfarisi/go-simple-crud/model"

type DepartmentRepositoryInterface interface {
	InsertDepartment(model.PostDepartment) bool
	//GetAllDepartment() []model.Department
	//GetOneDepartment(int) model.Department
	//DeleteDepartment(int) bool
	//UpdateDepartment(int, model.PostDepartment) model.Department
}
