package model

type Enroll struct {
	ID         int  `json:"id"`
	Course_Id  int  `json:"course_id"`
	Student_Id *int `json:"student_id"`
	Date       *int `json:"date"`
}

type PostEnroll struct {
	Course_Id  int  `json:"course_id"`
	Student_Id *int `json:"student_id"`
	Date       *int `json:"date"`
}

type EnrollUri struct {
	ID int `uri:"id" binding:"required,number"`
}
