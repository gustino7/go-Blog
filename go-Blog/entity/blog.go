package entity

type Blog struct {
	ID      uint64        `gorm:"primaryKey" json:"id"`
	Title   string        `json:"title"`
	Body    string        `json:"body"`
	Like    uint64        `json:"like"`
	Comment []BlogComment `json:"comment,omitempty"`
	UserID  uint64        `gorm:"foreignKey" json:"user_id"`
	User    *User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
}
