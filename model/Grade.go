package model

type Grade struct {
	ID         int  `json:"id"`
	Course_Id  int  `json:"course_id"`
	Student_Id *int `json:"student_id"`
	Score      *int `json:"score"`
}

type PostGrade struct {
	Course_Id  int  `json:"course_id"`
	Student_Id *int `json:"student_id"`
	Score      *int `json:"score"`
}

type GradeUri struct {
	ID int `uri:"id" binding:"required,number"`
}
