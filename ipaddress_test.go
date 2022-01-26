package marshallable

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPBasic(t *testing.T) {
	data := `{"IP":"192.168.1.4"}`
	payload := new(struct{ IP IP })
	err := json.Unmarshal([]byte(data), payload)
	if err != nil {
		t.Fatalf("IP failed to Unmarshal: %v", err)
	}

	o, err := json.Marshal(payload)
	t.Logf("payload: %s", o)
	assert.NoError(t, err)
	assert.Equal(t, string(o), data)
}

func TestIPNil(t *testing.T) {
	var IP IP
	V := IP.String()
	t.Logf("IP: %s", V)
}
