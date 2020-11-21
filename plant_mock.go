package ecologi

import (
	"encoding/json"
	"fmt"
)

// MockPlant provides a mock matching a successful POST to the /impact/trees
// endpoint. This allows users to test their applications using
// ecologi-client-go without incurring charges
// Assumptions:
// 1. the returned amount is 0.12 * number
// 2. the returned TreeURL uses mocked username `bento`, regardless of
// the owner and validity of the input API token.
func (c *Client) MockPlant(number int, name string, test bool) (*PlantOrderResponse, error) {
	order := PlantOrder{
		Number: number,
		Name:   name,
		Test:   test,
	}

	// verify the inputs can be marshalled correctly
	_, err := json.Marshal(order)

	if err != nil {
		return &PlantOrderResponse{}, fmt.Errorf("Failed to marshal order: %w", err)
	}

	// when testing, /impact/trees does not return JSON that can be
	// decoded to a valid PlantOrderResponse.
	if test {
		return &PlantOrderResponse{}, nil
	}

	// return a mocked response
	orderResponse := PlantOrderResponse{
		Amount:   float64(number) * defaultTreeValue,
		Currency: defaultCurrency,
		Name:     name,
		TreeURL:  "https://ecologi.com/" + defaultUsername + "?tree=000000000000000000000000",
	}

	return &orderResponse, nil

}
