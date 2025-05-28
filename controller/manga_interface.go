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

	//department
	InsertDepartment(*gin.Context)
	GetAllDepartment(*gin.Context)
	GetOneDepartment(*gin.Context)
	UpdateDepartment(*gin.Context)
	DeleteDepartment(*gin.Context)

	//student
	InsertStudent(*gin.Context)
	GetAllStudent(*gin.Context)
	GetOneStudent(*gin.Context)
	UpdateStudent(*gin.Context)
	DeleteStudent(*gin.Context)

	//enroll
	InsertEnroll(*gin.Context)
	GetAllEnroll(*gin.Context)
	GetOneEnroll(*gin.Context)
	UpdateEnroll(*gin.Context)
	DeleteEnroll(*gin.Context)

	//grade
	InsertGrade(*gin.Context)
	GetAllGrade(*gin.Context)
	GetOneGrade(*gin.Context)
	UpdateGrade(*gin.Context)
	DeleteGrade(*gin.Context)

	//teacher
	InsertTeacher(*gin.Context)
}

type ClassroomControllerInterface interface {
	//InsertClassroom(*gin.Context)
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

type DepartmentControllerInterface interface {
	//InsertDepartment(*gin.Context)
	//GetAllDepartment(*gin.Context)
	//GetOneDepartment(*gin.Context)
	//UpdateDepartment(*gin.Context)
	//DeleteDepartment(*gin.Context)
}
type StudentControllerInterface interface {
	//InsertStudent(*gin.Context)
	//GetAllStudent(*gin.Context)
	//GetOneStudent(*gin.Context)
	//UpdateStudent(*gin.Context)
	//DeleteStudent(*gin.Context)
}
type EnrollControllerInterface interface {
	//InsertEnroll(*gin.Context)
	//GetAllEnroll(*gin.Context)
	//GetOneEnroll(*gin.Context)
	//UpdateEnroll(*gin.Context)
	//DeleteEnroll(*gin.Context)
}
type GradeControllerInterface interface {
	//InsertGrade(*gin.Context)
	//GetAllGrade(*gin.Context)
	//GetOneGrade(*gin.Context)
	//UpdateGrade(*gin.Context)
	//DeleteGrade(*gin.Context)
}
type TeacherControllerInterface interface {
	//InsertTeacher(*gin.Context)
	//GetAllTeacher(*gin.Context)
	//GetOneTeacher(*gin.Context)
	//UpdateTeacher(*gin.Context)
	//DeleteTeacher(*gin.Context)
}
