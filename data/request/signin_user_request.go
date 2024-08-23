package request

type SigninRequest struct {
	Email    string `validate:"required,email,max=255" json:"email"`
	Password string `validate:"required,min=1,max=255" json:"password"`
}
