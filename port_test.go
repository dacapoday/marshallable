package marshallable

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPortBasic(t *testing.T) {
	data := `{"Port":8080}`
	payload := new(struct{ Port Port })
	err := json.Unmarshal([]byte(data), payload)
	if err != nil {
		t.Fatalf("Port failed to Unmarshal: %v", err)
	}

	assert.Equal(t, payload.Port.Port, uint16(8080))

	o, err := json.Marshal(payload)
	t.Logf("payload: %s", o)
	assert.NoError(t, err)
	assert.Equal(t, string(o), data)
	assert.Equal(t, payload.Port.String(), `8080`)
}

func TestPortNil(t *testing.T) {
	var Port Port
	V := Port.String()
	t.Logf("Port: %s", V)
	v, err := Port.MarshalJSON()
	assert.NoError(t, err)
	t.Logf("Port: %s", v)
}
