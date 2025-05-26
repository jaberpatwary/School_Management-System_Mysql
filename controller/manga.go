package controller

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ramadhanalfarisi/go-simple-crud/model"
	"github.com/ramadhanalfarisi/go-simple-crud/repository"
)

type MangaController struct {
	Db *sql.DB
}

func NewMangaController(db *sql.DB) MangaControllerInterface {
	return &MangaController{Db: db}
}

// insart classroom
func (m *MangaController) InsertClassroom(c *gin.Context) {

	DB := m.Db
	var post model.PostClassroom
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewClassroomRepository(DB)
	insert := repository.InsertClassroom(post)
	if insert {
		c.JSON(201, gin.H{"Status": "success", "meg": " classroom has saved!"})
		return
	} else {
		c.JSON(500, gin.H{"Status": "failed", "meg": " classrooom  is not saved!"})
		return
	}
}

// Gell all classroom
func (m *MangaController) GetAllClassroom(c *gin.Context) {
	DB := m.Db
	repository := repository.NewClassroomRepository(DB)
	get := repository.GetAllClassroom()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get classroom successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "classroom not found"})
		return
	}
}

// Get One classroom
func (m *MangaController) GetOneClassroom(c *gin.Context) {
	DB := m.Db
	var uri model.ClassroomUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewClassroomRepository(DB)
	get := repository.GetOneClassroom(uri.ID)
	if (get != model.Classroom{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get classroom successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "classroom not found"})
		return
	}
}

//update classroom

func (m *MangaController) UpdateClassroom(c *gin.Context) {
	DB := m.Db
	var post model.PostClassroom
	var uri model.ClassroomUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewClassroomRepository(DB)
	update := repository.UpdateClassroom(uri.ID, post)
	if (update != model.Classroom{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "classroom ": "update classroom successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "classroom ": "update classroom  failed"})
		return
	}
}

// Deleteclassroom
func (m *MangaController) DeleteClassroom(c *gin.Context) {
	DB := m.Db
	var uri model.ClassroomUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewClassroomRepository(DB)
	delete := repository.DeleteClassroom(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete classroom successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete classroom  failed"})
		return
	}
}

// insart course
func (m *MangaController) InsertCourse(c *gin.Context) {

	DB := m.Db
	var post model.PostCourse
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewCourseRepository(DB)
	insert := repository.InsertCourse(post)
	if insert {
		c.JSON(201, gin.H{"Status": "success", "meg": " course has saved!"})
		return
	} else {
		c.JSON(500, gin.H{"Status": "failed", "meg": " course  is not saved!"})
		return
	}
}

// Gell all course
func (m *MangaController) GetAllCourse(c *gin.Context) {
	DB := m.Db
	repository := repository.NewCourseRepository(DB)
	get := repository.GetAllCourse()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get course successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "course not found"})
		return
	}
}

// Get One course
func (m *MangaController) GetOneCourse(c *gin.Context) {
	DB := m.Db
	var uri model.CourseUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewCourseRepository(DB)
	get := repository.GetOneCourse(uri.ID)
	if (get != model.Course{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get course successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "course not found"})
		return
	}
}

//update course

func (m *MangaController) UpdateCourse(c *gin.Context) {
	DB := m.Db
	var post model.PostCourse
	var uri model.CourseUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewCourseRepository(DB)
	update := repository.UpdateCourse(uri.ID, post)
	if (update != model.Course{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "course ": "update course successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "course ": "update course failed"})
		return
	}
}

// Deletecourse
func (m *MangaController) DeleteCourse(c *gin.Context) {
	DB := m.Db
	var uri model.CourseUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewCourseRepository(DB)
	delete := repository.DeleteCourse(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete course successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete course  failed"})
		return
	}
}
