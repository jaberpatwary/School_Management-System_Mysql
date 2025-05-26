package model

type Student struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Date_Of_Birth *int    `json:"date_of_birth"`
	Email         *string `json:"email"`
}

type PostStudent struct {
	Name          string  `json:"name"`
	Date_Of_Birth *int    `json:"date_of_birth"`
	Email         *string `json:"email"`
}

type StudentUri struct {
	ID int `uri:"id" binding:"required,number"`
}
