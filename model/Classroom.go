package model

type classroom struct {
	ID       int  `json:"id"`
	Room     int  `json:"room"`
	Capacity *int `json:"capacity"`
}

type PostClassroom struct {
	Room     int  `json:"room"`
	Capacity *int `json:"capacity"`
}

type ClassroomUri struct {
	ID int `uri:"id" binding:"required,number"`
}
