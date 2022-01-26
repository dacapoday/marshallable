package marshallable

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlBasic(t *testing.T) {
	data := `{"Url":"http://example.com"}`
	payload := new(struct{ Url URL })
	err := json.Unmarshal([]byte(data), payload)
	if err != nil {
		t.Fatalf("URL failed to Unmarshal: %v", err)
	}

	o, err := json.Marshal(payload)
	t.Logf("payload: %s", o)
	assert.NoError(t, err)
	assert.Equal(t, string(o), data)
}

func TestUrlNil(t *testing.T) {
	var Url URL
	V := Url.String()
	t.Logf("Url: %s", V)
}
