package repository

import "github.com/ramadhanalfarisi/go-simple-crud/model"

type CategoryRepositoryInterface interface {
	InsertCategory(model.PostCategory) bool
	GetAllCategory() []model.Category
	GetOneCategory(int) model.Category
	DeleteCategory(int) bool
	UpdateCategory(int, model.PostCategory) model.Category
}
