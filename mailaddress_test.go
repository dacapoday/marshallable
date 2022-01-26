package marshallable

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMailAddressBasic(t *testing.T) {
	data := `{"MailAddress":"someone@example.com"}`
	payload := new(struct{ MailAddress MailAddress })
	err := json.Unmarshal([]byte(data), payload)
	if err != nil {
		t.Fatalf("MailAddress failed to Unmarshal: %v", err)
	}

	o, err := json.Marshal(payload)
	t.Logf("payload: %s", o)
	assert.NoError(t, err)
	assert.Equal(t, string(o), data)
}

func TestMailAddressNil(t *testing.T) {
	var MailAddress MailAddress
	V := MailAddress.String()
	t.Logf("MailAddress: %s", V)
}
