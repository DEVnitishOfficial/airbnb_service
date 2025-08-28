package middlewares

import (
	"AuthInGo/dto"
	"AuthInGo/utils"
	"net/http"
)

func UserLoginRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.LoginUserRequestDto // define the payload we want

		// Read and decode the json body into the payload
		if err := utils.ReadJSONBody(r, &payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}

		// validate the payload using the validator instance
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "validation failed", err)
			return
		}
		next.ServeHTTP(w, r) // call the next handler in the chain
	})
}

func UserCreateRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.CreateUserRequestDto // define the payload we want

		// Read and decode the json body into the payload
		if err := utils.ReadJSONBody(r, &payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}

		// validate the payload using the validator instance
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "validation failed", err)
			return
		}
		next.ServeHTTP(w, r) // call the next handler in the chain
	})
}
