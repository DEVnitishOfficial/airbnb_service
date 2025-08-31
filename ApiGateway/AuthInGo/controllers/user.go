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

func (uc *UserController) GetAllUserController(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GetAllUserService called from the controller layer")
	user, err := uc.userService.GetAllUserService()

	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch user", err)
		return
	}
	if user == nil {
		utils.WriteJSONErrorResponse(w, http.StatusNotFound, "User not found", err)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "User fetched successfully", user)
	fmt.Println("User fetched successfully:", user)

}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUserById is called in UserController")

	userId := r.URL.Query().Get("id")
	if userId == "" {
		userId = r.Context().Value("userID").(string)
	}
	fmt.Println("userId from the context or query param:", userId)

	if userId == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "User ID is required", fmt.Errorf("user ID is required"))
		return
	}

	user, err := uc.userService.GetUserById(userId)

	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch user", err)
		return
	}
	if user == nil {
		utils.WriteJSONErrorResponse(w, http.StatusNotFound, "User not found", fmt.Errorf("user with ID %d not found", userId))
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "User fetched successfully", user)
	fmt.Println("User fetched successfully:", user)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser is called in UserController")

	payload := r.Context().Value("payload").(dto.CreateUserRequestDto)

	user, err := uc.userService.CreateUser(&payload)

	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to create user", err)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "User created successfully", user)
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("LoginUser is called in UserController")

	payload := r.Context().Value("payload").(dto.LoginUserRequestDto)

	fmt.Println("LoginUserRequestDto:", payload)

	jwtToken, err := uc.userService.LoginUser(&payload)

	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to login user", err)
		return
	}

	// utils.WriteJSONResponse(w, http.StatusOK, response)
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "User loggedIn successfully", jwtToken)

}
