package ecologi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

// TestPlantInvalidToken ensures that, given a valid input to c.Plant but
// invalid API token, a 401 error is recieved
func TestPlantInvalidToken(t *testing.T) {
	token := defaultDummyAPIToken
	c, err := NewClient(nil, &token)

	assert.NoError(t, err)

	_, err = c.Plant(1, "Customer #123", false)

	// An invalid token should bring back an authentication error
	assert.Error(t, err)

	assert.EqualError(t, err, "POST request failed: failed to create resource with HTTP response code: 401")
}

// TestPlantMockResponse mocks a valid ecologi API response at /impact/trees
// as an easy way to test c.Plant makes a valid POST request to /impact/trees
// given valid imput parameters
func TestPlantMockResponse(t *testing.T) {

	expected := &PlantOrderResponse{
		Amount:   0.12,
		Currency: "USD",
		Name:     "Customer #123",
		TreeURL:  "https://ecologi.com/bento?tree=000000000000000000000000",
	}
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("https://public.ecologi.com").
		Post("/impact/trees").
		Reply(201).
		JSON(PlantOrderResponse{
			Amount:   0.12,
			Currency: "USD",
			Name:     "Customer #123",
			TreeURL:  "https://ecologi.com/bento?tree=000000000000000000000000",
		})

	token := defaultDummyAPIToken
	c, err := NewClient(nil, &token)

	gock.InterceptClient(c.HTTPClient)

	assert.NoError(t, err)

	resp, err := c.Plant(1, "Customer #123", false)

	assert.NoError(t, err)

	assert.Equal(t, expected, resp)

}
