package entity

type BlogComment struct {
	Comment string `json:"comment"`
	BlogID  uint64 `gorm:"foreignKey" json:"blog_id" binding:"required"`
	Blog    *Blog  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"blog,omitempty"`
}
