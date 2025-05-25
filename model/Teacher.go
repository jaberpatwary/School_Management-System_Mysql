package model

type Teacher struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Department_Id *int    `json:"department_id"`
	Email         *string `json:"Email"`
}

type PostTeacher struct {
	Name          string  `json:"name"`
	Department_Id *int    `json:"department_id"`
	Email         *string `json:"Email"`
}

type TeacherUri struct {
	ID int `uri:"id" binding:"required,number"`
}
