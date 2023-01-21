package mod_auth

type LoginValidation struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
