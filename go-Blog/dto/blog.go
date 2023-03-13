package dto

type UpBlog struct {
	ID     uint64 `gorm:"primaryKey" json:"id"`
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
	UserID uint64 `json:"user_id" binding:"required"`
}
