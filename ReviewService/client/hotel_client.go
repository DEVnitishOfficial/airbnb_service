package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HotelClient struct {
	BaseURL    string
	HttpClient *http.Client
}

func NewHotelClient(baseURL string) *HotelClient {
	return &HotelClient{
		BaseURL:    baseURL,
		HttpClient: &http.Client{},
	}
}

type HotelRating struct {
	Rating      json.Number `json:"rating"`
	RatingCount int64       `json:"ratingCount"`
}

// GetHotelRating fetches the current rating and rating_count for a given hotel ID
func (c *HotelClient) GetHotelRating(hotelID int64) (*HotelRating, error) {

	url := fmt.Sprintf("%s/hotels/id/%d", c.BaseURL, hotelID)

	fmt.Println("Fetching hotel rating from URL:", url)

	resp, err := c.HttpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("get hotel rating: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("hotel service returned %d", resp.StatusCode)
	}

	var res struct {
		Data HotelRating `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	// print fetched data
	fmt.Println("Fetched hotel rating data:", res.Data)

	return &res.Data, nil
}

// UpdateHotelRating updates the average rating and rating_count in the HotelService
func (c *HotelClient) UpdateHotelRating(hotelID int64, rating float64, count int64) error {
	url := fmt.Sprintf("%s/hotels/updateById/%d", c.BaseURL, hotelID)

	fmt.Println("Updating hotel rating at URL:", url)

	payload := map[string]interface{}{
		"rating":      rating,
		"ratingCount": count,
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("patch hotel rating: %w", err)
	}
	defer resp.Body.Close()

	fmt.Println("Hotel service response status:", resp.Status)
	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Println("Hotel service response data:", string(bodyBytes))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("hotel update failed: %d", resp.StatusCode)
	}
	return nil
}
