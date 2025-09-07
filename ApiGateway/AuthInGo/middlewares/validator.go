package middlewares

import (
	"AuthInGo/dto"
	"AuthInGo/utils"
	"context"
	"fmt"
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

		fmt.Println("UserLoginRequestValidator passed, payload:", payload)

		ctx := context.WithValue(r.Context(), "payload", payload)

		next.ServeHTTP(w, r.WithContext(ctx)) // call the next handler in the chain
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
		fmt.Println("payload reveived for login", payload)

		ctx := context.WithValue(r.Context(), "payload", payload) // create a new context with the payload

		next.ServeHTTP(w, r.WithContext(ctx)) // call the next handler in the chain
	})
}

func RoleCreateRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.CreateRoleRequestDto
		if err := utils.ReadJSONBody(r, &payload); err != nil {
			fmt.Println("error in reading json body", err)
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "validation failed", err)
			return
		}
		fmt.Println("payload reveived for create role", payload)
		ctx := context.WithValue(r.Context(), "payload", payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RoleUpdateRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.UpdateRoleRequestDto
		if err := utils.ReadJSONBody(r, &payload); err != nil {
			fmt.Println("error in reading json body", err)
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "validation failed", err)
			return
		}
		fmt.Println("payload reveived for update role", payload)
		ctx := context.WithValue(r.Context(), "payload", payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func PermissionCreateRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.CreatePermissionRequestDto
		if err := utils.ReadJSONBody(r, &payload); err != nil {
			fmt.Println("error in reading json body", err)
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "validation failed", err)
			return
		}
		fmt.Println("payload reveived for create Permission", payload)
		ctx := context.WithValue(r.Context(), "payload", payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func PermissionUpdateRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.UpdatePermissionRequestDto
		if err := utils.ReadJSONBody(r, &payload); err != nil {
			fmt.Println("error in reading json body", err)
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "validation failed", err)
			return
		}
		fmt.Println("payload reveived for update permission", payload)
		ctx := context.WithValue(r.Context(), "payload", payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
