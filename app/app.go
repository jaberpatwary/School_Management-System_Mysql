package app

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ramadhanalfarisi/go-simple-crud/controller"
	"github.com/ramadhanalfarisi/go-simple-crud/db"
)

type App struct {
	DB     *sql.DB
	Router *gin.Engine
}

func (a *App) CreateConnection() {
	db := db.Connectdb()
	a.DB = db
}

func (a *App) Routes() {
	r := gin.Default()
	controller := controller.NewMangaController(a.DB)

	//classroom router
	r.POST("/classroom", controller.InsertClassroom)
	r.GET("/classroom", controller.GetAllClassroom)
	r.GET("/classroom/:id", controller.GetOneClassroom)
	r.PUT("/classroom/:id", controller.UpdateClassroom)
	r.DELETE("/classroom/:id", controller.DeleteClassroom)

	//course
	r.POST("/course", controller.InsertCourse)
	r.GET("/course", controller.GetAllCourse)
	r.GET("/course/:id", controller.GetOneCourse)
	r.PUT("/course/:id", controller.UpdateCourse)
	r.DELETE("/course/:id", controller.DeleteCourse)

	//department
	r.POST("/department", controller.InsertDepartment)
	r.GET("/department", controller.GetAllDepartment)
	r.GET("/department/:id", controller.GetOneDepartment)
	r.PUT("/department/:id", controller.UpdateDepartment)
	r.DELETE("/department/:id", controller.DeleteDepartment)

	//student
	r.POST("/student", controller.InsertStudent)
	r.GET("/student", controller.GetAllStudent)
	r.GET("/student/:id", controller.GetOneStudent)
	r.PUT("/student/:id", controller.UpdateStudent)
	r.DELETE("/student/:id", controller.DeleteStudent)

	//enroll
	r.POST("/enroll", controller.InsertEnroll)
	r.GET("/enroll", controller.GetAllEnroll)
	r.GET("/enroll/:id", controller.GetOneEnroll)
	r.PUT("/enroll/:id", controller.UpdateEnroll)
	r.DELETE("/enroll/:id", controller.DeleteEnroll)

	//grade
	r.POST("/grade", controller.InsertGrade)
	r.GET("/grade", controller.GetAllGrade)
	r.GET("/grade/:id", controller.GetOneGrade)
	r.PUT("/grade/:id", controller.UpdateGrade)
	r.DELETE("/grade/:id", controller.DeleteGrade)

	//teacher
	r.POST("/teacher", controller.InsertTeacher)
	r.GET("/teacher", controller.GetAllTeacher)
	r.GET("/teacher/:id", controller.GetOneTeacher)
	r.PUT("/teacher/:id", controller.UpdateTeacher)
	r.DELETE("/teacher/:id", controller.DeleteTeacher)

	a.Router = r
}

func (a *App) Run() {
	a.Router.Run(":2001")
}
