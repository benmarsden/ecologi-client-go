package ecologi

// PlantOrder is a Go-native representation of the JSON encoding for a POST
// request to the /impact/trees endpoint
type PlantOrder struct {
	Number int    `json:"number"`
	Name   string `json:"name,omitempty"`
	Test   bool   `json:"test,omitempty"`
}

// PlantOrderResponse is a Go-native representation of the JSON decoding from a
// successful response to the /impact/trees endpoint
type PlantOrderResponse struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Name     string  `json:"name"`
	TreeURL  string  `json:"treeURL"`
}

// TreeCount is a Go-native representation of the JSON decoding from a
// successful response to the /users/<username>/trees endpoint
type TreeCount struct {
	Total int `json:"total"`
}

// CarbonOffsetCount is a Go-native representation of the JSON decoding from a
// successful response to the /users/<username>/carbon-offset endpoint
type CarbonOffsetCount struct {
	Total int `json:"total"`
}

// TreeAndCarbonOffsetCount is a Go-native representation of the JSON decoding
// from a successful response to the /users/<username>/carbon-offset endpoint
type TreeAndCarbonOffsetCount struct {
	TreeTotal         int `json:"trees"`
	CarbonOffsetTotal int `json:"carbonOffset"`
}
