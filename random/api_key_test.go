package random_test

import (
	"testing"

	"github.com/ecumenos-social/toolkit/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenAPIKey(t *testing.T) {
	apiKey, err := random.GenAPIKey("nw", "1234567890")
	require.NoError(t, err)
	assert.Contains(t, apiKey, "nw")
	assert.Contains(t, apiKey, ".")
}
