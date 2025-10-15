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
/**
Is all the UserService property we are putting here inside the _userService and further
inside the userService in return statement?

Yes, that is exactly what is happening.

When you call NewUserController, you are required to pass in an object that implements the
services.UserService interface.

That object is received inside the NewUserController function and is given the local
name _userService.

The return statement creates a new UserController struct.

The line userService: _userService, takes the object you passed in (the value of
the _userService parameter) and assigns it to the userService field of the new
 UserController struct.

This is the standard way to inject a dependency.
**/
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

func (uc *UserController) GetBulkUserInfoByIdsController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetBulkUserInfoByIdsController is called in UserController")

	// Extract user IDs from the request body (as JSON array)
	var userIds []int64
	err := utils.ReadJSONBody(r, &userIds)

	fmt.Println("User IDs received in controller:", userIds)

	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	users, err := uc.userService.GetBulkUserInfoByIdsService(userIds)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch users", err)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Users fetched successfully", users)
}
