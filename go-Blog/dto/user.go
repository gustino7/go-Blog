package dto

type RegisterUser struct {
	Nama     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" type:"varchar(100)" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" type:"varchar(100)" binding:"required"`
}

type UserUpdateDTO struct {
	Nama string `json:"name" binding:"required"`
}
