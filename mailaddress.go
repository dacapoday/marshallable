package marshallable

import (
	"encoding/json"
	"net/mail"
)

// MailAddress is a mail.Address wrapper that marshals to and from a JSON string.
type MailAddress struct {
	*mail.Address
}

// String returns the MailAddress as a string.
func (m MailAddress) String() string {
	if m.Address == nil {
		return ""
	}
	return m.Address.Address
}

// MarshalJSON implements the json.Marshaler interface.
func (m MailAddress) MarshalJSON() ([]byte, error) {
	if m.Address == nil {
		return []byte("null"), nil
	}
	return json.Marshal(m.Address.Address)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (m *MailAddress) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
	if err != nil {
		return
	}
	m.Address, err = mail.ParseAddress(value)
	return
}

// MustMailAddress returns the MailAddress or panics.
func MustMailAddress(value string) MailAddress {
	mailAddress, err := mail.ParseAddress(value)
	if err != nil {
		panic(err)
	}
	return MailAddress{mailAddress}
}
