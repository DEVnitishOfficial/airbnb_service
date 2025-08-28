package dto

// Marshalling with validator
// Generally we define struct property first capital letter so that it is publically abailable
// and can access inside any package
// But when we send the jsone we write it in small like {"email":"nkum@62@gmail.com"}
// so simply we can tell our decoder whenever anyone send data like ({"email","password"}) then
// you have to map with sturct property like Email, Password etc.
// and this is done using the concept of json Marshalling.
type LoginUserRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type CreateUserRequestDto struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
