package requests

type UserRequest struct {
	Email    string `gorm:"unique" json:"email" validate:"email"`
	Password string `json:"password" validate:"min=4"`
}
