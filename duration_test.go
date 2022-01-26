package marshallable

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDurationBasic(t *testing.T) {
	data := `{"Duration":"1s"}`
	payload := new(struct{ Duration Duration })
	err := json.Unmarshal([]byte(data), payload)
	if err != nil {
		t.Fatalf("Duration failed to Unmarshal: %v", err)
	}

	duration1s, err := time.ParseDuration("1s")
	assert.NoError(t, err)
	assert.Equal(t, payload.Duration.Duration, duration1s)

	o, err := json.Marshal(payload)
	t.Logf("payload: %s", o)
	assert.NoError(t, err)
	assert.Equal(t, string(o), data)
}

func TestDurationNil(t *testing.T) {
	var Duration Duration
	V := Duration.String()
	t.Logf("Duration: %s", V)
}
