package ecologi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPlantMock - tests PlantMock
func TestPlantMock(t *testing.T) {
	token := defaultDummyAPIToken
	c, err := NewClient(nil, &token)

	assert.NoError(t, err)

	orderRes, err := c.MockPlant(3, "Customer #123", false)

	assert.NoError(t, err)

	t.Logf("Response: %+v", orderRes)

	assert.NotNil(t, orderRes)

	assert.Equal(t, 3*defaultTreeValue, orderRes.Amount)
}
