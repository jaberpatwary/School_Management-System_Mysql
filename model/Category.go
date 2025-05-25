package model

type Category struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`  // Required field
	Types *string `json:"types"` // Required field

}

type PostCategory struct {
	Name  string  `json:"name"` // Required field
	Types *string `json:"types` // Required field

}

type CategoryUri struct {
	ID int `uri:"id" binding:"required,number"`
}
