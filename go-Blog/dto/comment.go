package dto

type AddComm struct {
	Title   string `json:"title" binding:"required"`
	Comment string `json:"comment" binding:"required"`
}
