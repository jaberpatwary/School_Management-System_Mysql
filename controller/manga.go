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

/*Insertcategory

func (m *MangaController) InsertCategory(c *gin.Context) {

	DB := m.Db
	var post model.PostCategory
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewCategoryRepository(DB)
	insert := repository.InsertCategory(post)
	if insert {
		c.JSON(201, gin.H{"Status": "success", "meg": " category has saved!"})
		return
	} else {
		c.JSON(500, gin.H{"Status": "failed", "meg": " category  is not saved!"})
		return
	}

}

// Gell all category
func (m *MangaController) GetAllCategory(c *gin.Context) {
	DB := m.Db
	repository := repository.NewCategoryRepository(DB)
	get := repository.GetAllCategory()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get category successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "category not found"})
		return
	}
}

// Get One cagetory
func (m *MangaController) GetOneCategory(c *gin.Context) {
	DB := m.Db
	var uri model.CategoryUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewCategoryRepository(DB)
	get := repository.GetOneCategory(uri.ID)
	if (get != model.Category{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get category successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "category not found"})
		return
	}
}

//update category

func (m *MangaController) UpdateCategory(c *gin.Context) {
	DB := m.Db
	var post model.PostCategory
	var uri model.CategoryUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewCategoryRepository(DB)
	update := repository.UpdateCategory(uri.ID, post)
	if (update != model.Category{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "category": "update category successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "category": "update category failed"})
		return
	}
}

//Deletecategory

func (m *MangaController) DeleteCategory(c *gin.Context) {
	DB := m.Db
	var uri model.CategoryUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewCategoryRepository(DB)
	delete := repository.DeleteCategory(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete category successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete category failed"})
		return
	}
}*/
//Insertcategory

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
