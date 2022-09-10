package request

type UserRegisterReq struct {
	Name     string `json:"name" gorm:"name" validate:"required"`
	Email    string `json:"email" gorm:"email" validate:"required,email"`
	Password string `json:"password" gorm:"password" validate:"required,min=6"`
}

type UserLoginReq struct {
	Email    string `json:"email" gorm:"email" validate:"required,email"`
	Password string `json:"password" gorm:"password" validate:"required,min=6"`
}
