package model

type Course struct {
	ID          int     `json:"id"`
	Course_Name string  `json:"course_name"`
	Department  *string `json:"department"`
	Credit      *string `json:"credit"`
}

type PostCourse struct {
	Course_Name string  `json:"course_name"`
	Department  *string `json:"department"`
	Credit      *string `json:"credit"`
}

type CourseUri struct {
	ID int `uri:"id" binding:"required,number"`
}
