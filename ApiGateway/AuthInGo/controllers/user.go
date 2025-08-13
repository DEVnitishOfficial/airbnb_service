// Now here similir to the userService which depend on the userRepository
// controller is dependent on the userService
// So we will use the similir approach(interface) to create a controller

package controllers

import (
	services "AuthInGo/services"
	"fmt"
	"net/http"
)

// this controller will be dependent on the userService
type UserController struct {
	userService services.UserService // dependency injection of UserService
}

// constructor for UserController
func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		userService: _userService,
	}
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUserById is called in UserController")
	// calling userService to create a user
	uc.userService.GetUserById()
	w.Write([]byte("User fetching endpoint done"))
}
