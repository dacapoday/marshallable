package marshallable

import (
	"encoding/json"
	"regexp"
)

// Regexp is a regexp.Regexp wraper that marshals to and from a JSON string.
type Regexp struct{ *regexp.Regexp }

// String returns the Regexp as a string.
func (r Regexp) String() string {
	if r.Regexp == nil {
		return ""
	}
	return r.Regexp.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (r Regexp) MarshalJSON() ([]byte, error) {
	if r.Regexp == nil {
		return []byte("null"), nil
	}
	return json.Marshal(r.Regexp.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (r *Regexp) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
	if err != nil {
		return
	}
	r.Regexp, err = regexp.Compile(value)
	return
}

// MustRegexp returns the Regexp or panics.
func MustRegexp(value string) Regexp {
	regexp, err := regexp.Compile(value)
	if err != nil {
		panic(err)
	}
	return Regexp{regexp}
}

// POSIXRegexp is a POSIX regexp.Regexp wraper that marshals to and from a JSON string.
type POSIXRegexp struct{ *regexp.Regexp }

// String returns the POSIXRegexp as a string.
func (r POSIXRegexp) String() string {
	if r.Regexp == nil {
		return ""
	}
	return r.Regexp.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (r POSIXRegexp) MarshalJSON() ([]byte, error) {
	if r.Regexp == nil {
		return []byte("null"), nil
	}
	return json.Marshal(r.Regexp.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (r *POSIXRegexp) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
	if err != nil {
		return
	}
	r.Regexp, err = regexp.CompilePOSIX(value)
	return
}

// MustPOSIXRegexp returns the POSIXRegexp or panics.
func MustPOSIXRegexp(value string) POSIXRegexp {
	regexp, err := regexp.CompilePOSIX(value)
	if err != nil {
		panic(err)
	}
	return POSIXRegexp{regexp}
}
