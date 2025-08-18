// Now here similir to the userService which depend on the userRepository
// controller is dependent on the userService
// So we will use the similir approach(interface) to create a controller

package controllers

import (
	"AuthInGo/dto"
	services "AuthInGo/services"
	"AuthInGo/utils"
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

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser is called in UserController")
	res := uc.userService.CreateUser()
	fmt.Println("created user", res)
	w.Write([]byte("CreateUser endpoint done "))
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("LoginUser is called in UserController")

	var payload dto.LoginUserRequestDto

	if JsonErr := utils.ReadJSONBody(r, &payload); JsonErr != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Something went wrong while reading ReadJSONBody from controllers", JsonErr)
	}

	if ValidationErr := utils.Validator.Struct(payload); ValidationErr != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid input detail", ValidationErr)
		return
	}

	jwtToken, err := uc.userService.LoginUser(&payload)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to login user", err)
		return
	}

	// utils.WriteJSONResponse(w, http.StatusOK, response)
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "User loggedIn successfully", jwtToken)

}
