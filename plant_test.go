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
		TreeURL:  "https://ecologi.com/" + defaultTestUsername + "?tree=000000000000000000000000",
	}
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("https://public.ecologi.com").
		Post("/impact/trees").
		Reply(201).
		JSON(PlantOrderResponse{
			Amount:   0.12,
			Currency: "USD",
			Name:     "Customer #123",
			TreeURL:  "https://ecologi.com/" + defaultTestUsername + "?tree=000000000000000000000000",
		})

	token := defaultDummyAPIToken
	c, err := NewClient(nil, &token)

	gock.InterceptClient(c.HTTPClient)

	assert.NoError(t, err)

	resp, err := c.Plant(1, "Customer #123", false)

	assert.NoError(t, err)

	assert.Equal(t, expected, resp)

}

// TestGetTrees ensures that, given a valid username to c.GetTrees
// with known tree count > 0, a tree count greater than 0 is returned by the
// GET request
func TestGetTrees(t *testing.T) {
	c, err := NewClient(nil, nil)

	assert.NoError(t, err)

	tc, err := c.GetTrees(defaultTestUsername)

	assert.NoError(t, err)

	assert.Greater(t, tc.Total, 0)
}

// TestGetCarbonOffset ensures that, given a valid username to c.GetCarbonOffset
// with known carbon offset count > 0, a carbon offset count greater than 0 is
// returned by the GET request
func TestGetCarbonOffset(t *testing.T) {
	c, err := NewClient(nil, nil)

	assert.NoError(t, err)

	co, err := c.GetCarbonOffset(defaultTestUsername)

	assert.NoError(t, err)

	assert.Greater(t, co.Total, 0)
}

// TestGetTreesAndCarbonOffset ensures that, given a valid username to
// c.GetCarbonOffset with known carbon offset count > 0, carbon offset and tree
// counts greater than 0 are returned by the GET request
func TestGetTreesAndCarbonOffset(t *testing.T) {
	c, err := NewClient(nil, nil)

	assert.NoError(t, err)

	impact, err := c.GetTreesAndCarbonOffset(defaultTestUsername)

	assert.NoError(t, err)

	assert.Greater(t, impact.TreeTotal, 0)
	assert.Greater(t, impact.CarbonOffsetTotal, 0)
}
