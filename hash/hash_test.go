package hash_test

import (
	"testing"

	"github.com/ecumenos-social/toolkit/hash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	raw := "1234567890"
	h, err := hash.Hash(raw)
	require.NoError(t, err)
	assert.True(t, hash.CompareHash(raw, h))
}

func TestSHA256(t *testing.T) {
	raw := "1234567890"
	assert.Equal(t, hash.SHA256(raw), hash.SHA256(raw))
}

func TestSHA512(t *testing.T) {
	raw := "1234567890"
	assert.Equal(t, hash.SHA512(raw), hash.SHA512(raw))
}

func TestSHA1(t *testing.T) {
	raw := "1234567890"
	assert.Equal(t, hash.SHA1(raw), hash.SHA1(raw))
}

func TestMD5(t *testing.T) {
	raw := "1234567890"
	h1, err := hash.MD5(raw)
	require.NoError(t, err)

	h2, err := hash.MD5(raw)
	require.NoError(t, err)

	assert.Equal(t, h1, h2)
}

func TestFNV(t *testing.T) {
	raw := "1234567890"
	assert.Equal(t, hash.FNV(raw), hash.FNV(raw))
}

func TestAdler32(t *testing.T) {
	raw := "1234567890"
	assert.Equal(t, hash.Adler32(raw), hash.Adler32(raw))
}

func TestCRC32(t *testing.T) {
	raw := "1234567890"
	assert.Equal(t, hash.CRC32(raw), hash.CRC32(raw))
}

func TestCRC64(t *testing.T) {
	raw := "1234567890"
	assert.Equal(t, hash.CRC64(raw), hash.CRC64(raw))
}
