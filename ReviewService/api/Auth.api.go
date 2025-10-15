package api

import (
	"ReviewService/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchUsersByIds(userIds []int64, authHeader string) ([]dto.UserDTO, error) {
	url := "http://localhost:3002/getbulkinfobyids"

	jsonData, err := json.Marshal(userIds)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal userIds: %w", err)
	}

	fmt.Println("Request payload for user service:", string(jsonData))
	fmt.Println("Auth Header being sent to user service:", authHeader)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	if authHeader != "" {
		req.Header.Set("Authorization", authHeader)
		fmt.Println("Authorization header added:", authHeader)
	} else {
		fmt.Println("No Authorization header provided")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("non-200 response: %s", string(bodyBytes))
	}

	var respData struct {
		Data map[string]dto.UserDTO `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	var users []dto.UserDTO
	for _, user := range respData.Data {
		users = append(users, user)
	}

	return users, nil
}
