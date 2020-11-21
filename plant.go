package ecologi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

const (
	defaultTreeValue     float64 = 0.12
	defaultCurrency      string  = "USD"
	defaultTestUsername  string  = "bento"
	defaultDummyAPIToken string  = "dummy-token"
)

// Plant matches the POST /impact/trees endpoint, allowing you to purchase 1
// or more trees per request
// If successful: returns a non-nil PlantOrderResponse
// If testing: returns a nil PlantOrderResponse
// If unsuccessful: returns a nil PlantOrderResponse and error
func (c *Client) Plant(number int, name string, test bool) (*PlantOrderResponse, error) {
	order := PlantOrder{
		Number: number,
		Name:   name,
		Test:   test,
	}
	orderBytes, err := json.Marshal(order)

	if err != nil {
		return &PlantOrderResponse{}, fmt.Errorf("Failed to marshal order: %w", err)
	}

	orderReq, err := http.NewRequest("POST", fmt.Sprintf("%s/impact/trees", c.HostURL), strings.NewReader(string(orderBytes)))
	if err != nil {
		return &PlantOrderResponse{}, fmt.Errorf("Failed to form POST request for order to ecologi: %w", err)
	}

	var bearer string = "Bearer " + c.Token

	orderReq.Header.Set("Authorization", bearer)
	orderReq.Header.Set("Content-Type", "application/json")
	// TODO: Support idempotency keys
	// req.Header.Set("Idempotency-Key", fmt.Sprint(idempKey))

	body, err := c.doRequest(orderReq)
	if err != nil {
		return &PlantOrderResponse{}, fmt.Errorf("POST request failed: %w", err)
	}

	// when testing, /impact/trees does not return JSON that can be
	// decoded to a valid PlantOrderResponse.
	if test {
		return &PlantOrderResponse{}, nil
	}

	orderResponse := PlantOrderResponse{}

	if err := json.NewDecoder(body).Decode(&orderResponse); err != nil {
		return &PlantOrderResponse{}, fmt.Errorf("failed to decode JSON: %w", err)
	}

	if reflect.DeepEqual(orderResponse, PlantOrderResponse{}) {
		return &PlantOrderResponse{}, fmt.Errorf("Failed to create ecologi order: %w", errors.New("ecologi PlantOrderResponse is nil"))
	}

	return &orderResponse, nil

}

type TreeCount struct {
	Total int `json:"total"`
}

// GetTrees matches the GET /users/<username>/trees endpoint, allowing you to
// access the number of trees associated with a particular user
// If successful: returns a non-nil TreeCount
// If unsuccessful: returns a nil TreeCount and error
func (c *Client) GetTrees(username string) (*TreeCount, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/%s/trees", c.HostURL, username), nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to form GET request for requesting trees associated with user %s: %w", username, err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to make GET request for requesting trees associated with user %s: %w", username, err)
	}

	treeResponse := TreeCount{}
	if err := json.NewDecoder(body).Decode(&treeResponse); err != nil {
		return &TreeCount{}, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return &treeResponse, nil
}
