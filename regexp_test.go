package marshallable

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexpBasic(t *testing.T) {
	data := `{"Regexp":".[123]+"}`
	payload := new(struct{ Regexp Regexp })
	err := json.Unmarshal([]byte(data), payload)
	if err != nil {
		t.Fatalf("Regexp failed to Unmarshal: %v", err)
	}

	o, err := json.Marshal(payload)
	t.Logf("payload: %s", o)
	assert.NoError(t, err)
	assert.Equal(t, string(o), data)
}

func TestPOSIXRegexpBasic(t *testing.T) {
	data := `{"POSIXRegexp":".[123]+"}`
	payload := new(struct{ POSIXRegexp POSIXRegexp })
	err := json.Unmarshal([]byte(data), payload)
	if err != nil {
		t.Fatalf("POSIXRegexp failed to Unmarshal: %v", err)
	}

	o, err := json.Marshal(payload)
	t.Logf("payload: %s", o)
	assert.NoError(t, err)
	assert.Equal(t, string(o), data)
}

func TestRegexpNil(t *testing.T) {
	var Regexp Regexp
	V := Regexp.String()
	t.Logf("Regexp: %s", V)
	var POSIXRegexp POSIXRegexp
	V = POSIXRegexp.String()
	t.Logf("POSIXRegexp: %s", V)
}
