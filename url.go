package marshallable

import (
	"encoding/json"
	"net/url"
)

// URL is a net/url.URL wraper that marshals to and from a JSON string.
type URL struct{ *url.URL }

// String returns the URL as a string.
func (u URL) String() string {
	if u.URL == nil {
		return ""
	}
	return u.URL.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (u URL) MarshalJSON() ([]byte, error) {
	if u.URL == nil {
		return []byte("null"), nil
	}
	return json.Marshal(u.URL.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (u *URL) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
	if err != nil {
		return
	}
	u.URL, err = url.Parse(value)
	return
}

// MustURL returns the URL or panics.
func MustURL(value string) URL {
	url, err := url.Parse(value)
	if err != nil {
		panic(err)
	}
	return URL{url}
}
