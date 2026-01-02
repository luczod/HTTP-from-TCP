package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeaderParse(t *testing.T) {
	// Test: Valid single header
	headers := NewHeaders()
	data := []byte("Host: localhost:42069\r\nFooFoo:          barbar\r\n\r\n")
	n, done, err := headers.Parse(data)

	require.NoError(t, err)
	require.NotNil(t, headers)
	host, _ := headers.Get("Host")
	assert.Equal(t, "localhost:42069", host)
	foo, _ := headers.Get("FooFoo")
	assert.Equal(t, "barbar", foo)
	empty, _ := headers.Get("MissingKey")
	assert.Equal(t, "", empty)
	assert.Equal(t, 50, n)
	assert.True(t, done)

	// Test: Invalid spacing header
	headers = NewHeaders()
	data = []byte("        Host  : localhost:42069            \r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	// Test: Invalid header name
	headers = NewHeaders()
	data = []byte("H@st:localhost:42069\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	// Test: Invalid header multi-values
	headers = NewHeaders()
	data = []byte("Host:localhost:42069\r\nHost: domain.com\r\n\r\n")
	n, done, err = headers.Parse(data)

	require.NoError(t, err)
	require.NotNil(t, headers)
	host, _ = headers.Get("Host")
	assert.Equal(t, "localhost:42069,domain.com", host)
}
