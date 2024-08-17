package request

type UpdateUserRequest struct {
	Id       int    `validate:"required"`
	FullName string `validate:"required,min=1,max=255" json:"fullname"`
	Password string `validate:"required,min=1,max=255" json:"password"`
}
