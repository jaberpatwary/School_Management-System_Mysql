package controller

import (
	"github.com/gin-gonic/gin"
)

type MangaControllerInterface interface {
	//classroom
	InsertClassroom(*gin.Context)
	GetAllClassroom(*gin.Context)
	GetOneClassroom(*gin.Context)
	UpdateClassroom(*gin.Context)
	DeleteClassroom(*gin.Context)

	//course
	InsertCourse(*gin.Context)
	GetAllCourse(*gin.Context)
	GetOneCourse(*gin.Context)
	UpdateCourse(*gin.Context)
	DeleteCourse(*gin.Context)
}

type ClassroomControllerInterface interface {
	//InsertClassroom(*gin.Context)
	//GetAllClassroom(*gin.Context)
	//GetOneClassroom(*gin.Context)
	//UpdateClassroom(*gin.Context)
	//DeleteClassroom(*gin.Context)
}

type CourseControllerInterface interface {
	//InsertCourse(*gin.Context)
	//GetAllCourse(*gin.Context)
	//GetOneCourse(*gin.Context)
	//UpdateCourse(*gin.Context)
	//DeleteCourse(*gin.Context)
}
