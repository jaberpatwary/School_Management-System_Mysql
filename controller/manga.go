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

// insart department
func (m *MangaController) InsertDepartment(c *gin.Context) {

	DB := m.Db
	var post model.PostDepartment
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewDepartmentRepository(DB)
	insert := repository.InsertDepartment(post)
	if insert {
		c.JSON(201, gin.H{"Status": "success", "meg": " department has saved!"})
		return
	} else {
		c.JSON(500, gin.H{"Status": "failed", "meg": " department  is not saved!"})
		return
	}
}

// Gell all department
func (m *MangaController) GetAllDepartment(c *gin.Context) {
	DB := m.Db
	repository := repository.NewDepartmentRepository(DB)
	get := repository.GetAllDepartment()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get department successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "department not found"})
		return
	}
}

// Get One departmnet
func (m *MangaController) GetOneDepartment(c *gin.Context) {
	DB := m.Db
	var uri model.DepartmentUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewDepartmentRepository(DB)
	get := repository.GetOneDepartment(uri.ID)
	if (get != model.Department{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get department successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "department not found"})
		return
	}
}

//update department

func (m *MangaController) UpdateDepartment(c *gin.Context) {
	DB := m.Db
	var post model.PostDepartment
	var uri model.DepartmentUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewDepartmentRepository(DB)
	update := repository.UpdateDepartment(uri.ID, post)
	if (update != model.Department{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "department ": "update department successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "department  ": "update department  failed"})
		return
	}
}

// Delete department
func (m *MangaController) DeleteDepartment(c *gin.Context) {
	DB := m.Db
	var uri model.DepartmentUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewDepartmentRepository(DB)
	delete := repository.DeleteDepartment(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete department successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete department  failed"})
		return
	}
}

// insart student
func (m *MangaController) InsertStudent(c *gin.Context) {

	DB := m.Db
	var post model.PostStudent
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewStudentRepository(DB)
	insert := repository.InsertStudent(post)
	if insert {
		c.JSON(201, gin.H{"Status": "success", "meg": " student has saved!"})
		return
	} else {
		c.JSON(500, gin.H{"Status": "failed", "meg": " student  is not saved!"})
		return
	}
}

// Gell all student
func (m *MangaController) GetAllStudent(c *gin.Context) {
	DB := m.Db
	repository := repository.NewStudentRepository(DB)
	get := repository.GetAllStudent()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get student successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "student not found"})
		return
	}
}

// Get One student
func (m *MangaController) GetOneStudent(c *gin.Context) {
	DB := m.Db
	var uri model.StudentUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewStudentRepository(DB)
	get := repository.GetOneStudent(uri.ID)
	if (get != model.Student{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get student successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "student not found"})
		return
	}
}

//update student

func (m *MangaController) UpdateStudent(c *gin.Context) {
	DB := m.Db
	var post model.PostStudent
	var uri model.StudentUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewStudentRepository(DB)
	update := repository.UpdateStudent(uri.ID, post)
	if (update != model.Student{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "student ": "update student successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "student  ": "update student  failed"})
		return
	}
}

// Delete student
func (m *MangaController) DeleteStudent(c *gin.Context) {
	DB := m.Db
	var uri model.StudentUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewStudentRepository(DB)
	delete := repository.DeleteStudent(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete student successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete student  failed"})
		return
	}
}

// insart enroll
func (m *MangaController) InsertEnroll(c *gin.Context) {

	DB := m.Db
	var post model.PostEnroll
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewEnrollRepository(DB)
	insert := repository.InsertEnroll(post)
	if insert {
		c.JSON(201, gin.H{"Status": "success", "meg": " enroll has saved!"})
		return
	} else {
		c.JSON(500, gin.H{"Status": "failed", "meg": " enroll  is not saved!"})
		return
	}
}

// Gell all enroll
func (m *MangaController) GetAllEnroll(c *gin.Context) {
	DB := m.Db
	repository := repository.NewEnrollRepository(DB)
	get := repository.GetAllEnroll()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get enroll successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "enroll not found"})
		return
	}
}

// Get One enroll
func (m *MangaController) GetOneEnroll(c *gin.Context) {
	DB := m.Db
	var uri model.EnrollUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewEnrollRepository(DB)
	get := repository.GetOneEnroll(uri.ID)
	if (get != model.Enroll{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get enroll successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "enroll not found"})
		return
	}
}

//update enroll

func (m *MangaController) UpdateEnroll(c *gin.Context) {
	DB := m.Db
	var post model.PostEnroll
	var uri model.EnrollUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewEnrollRepository(DB)
	update := repository.UpdateEnroll(uri.ID, post)
	if (update != model.Enroll{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "enroll ": "update enroll successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "enroll  ": "update enroll  failed"})
		return
	}
}

// Delete enroll
func (m *MangaController) DeleteEnroll(c *gin.Context) {
	DB := m.Db
	var uri model.EnrollUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewEnrollRepository(DB)
	delete := repository.DeleteEnroll(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete enroll successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete enroll  failed"})
		return
	}
}

// insart grade
func (m *MangaController) InsertGrade(c *gin.Context) {

	DB := m.Db
	var post model.PostGrade
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewGradeRepository(DB)
	insert := repository.InsertGrade(post)
	if insert {
		c.JSON(201, gin.H{"Status": "success", "meg": " grade has saved!"})
		return
	} else {
		c.JSON(500, gin.H{"Status": "failed", "meg": " grade  is not saved!"})
		return
	}
}

// Gell all grade
func (m *MangaController) GetAllGrade(c *gin.Context) {
	DB := m.Db
	repository := repository.NewGradeRepository(DB)
	get := repository.GetAllGrade()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get grade successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "grade not found"})
		return
	}
}

// Get One grade
func (m *MangaController) GetOneGrade(c *gin.Context) {
	DB := m.Db
	var uri model.GradeUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewGradeRepository(DB)
	get := repository.GetOneGrade(uri.ID)
	if (get != model.Grade{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get grade successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "grade not found"})
		return
	}
}

// update grade
func (m *MangaController) UpdateGrade(c *gin.Context) {
	DB := m.Db
	var post model.PostGrade
	var uri model.GradeUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewGradeRepository(DB)
	update := repository.UpdateGrade(uri.ID, post)
	if (update != model.Grade{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "grade ": "update grade successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "grade  ": "update grade failed"})
		return
	}
}

// Delete grade
func (m *MangaController) DeleteGrade(c *gin.Context) {
	DB := m.Db
	var uri model.GradeUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewGradeRepository(DB)
	delete := repository.DeleteGrade(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete grade successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete grade  failed"})
		return
	}
}

// insart teacher
func (m *MangaController) InsertTeacher(c *gin.Context) {

	DB := m.Db
	var post model.PostTeacher
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewTeacherRepository(DB)
	insert := repository.InsertTeacher(post)
	if insert {
		c.JSON(201, gin.H{"Status": "success", "meg": " teacher has saved!"})
		return
	} else {
		c.JSON(500, gin.H{"Status": "failed", "meg": " teacher  is not saved!"})
		return
	}
}
