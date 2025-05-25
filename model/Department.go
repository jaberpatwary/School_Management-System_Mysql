package model

type Department struct {
	ID       int  `json:"id"`
	Name     string  `json:"name"`
	Department_Head *string `json:"department_head"`
}

type PostDepartmentstruct {
	Name     string  `json:"name"`
	Department_Head *string `json:"department_head"`
}

type DepartmentUri struct {
	ID int `uri:"id" binding:"required,number"`
}
