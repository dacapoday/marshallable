package marshallable

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Port is a TCP/UDP port number that marshals to and from a JSON string or number.
type Port struct{ Port uint16 }

// String returns the port number as a string.
func (p Port) String() string {
	return strconv.FormatUint(uint64(p.Port), 10)
}

// MarshalJSON implements the json.Marshaler interface.
func (p Port) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Port)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (p *Port) UnmarshalJSON(b []byte) (err error) {
	var v interface{}
	if err = json.Unmarshal(b, &v); err != nil {
		return
	}
	switch value := v.(type) {
	case float64:
		if value < 0 || value > 65535 {
			err = fmt.Errorf("invalid port number: %v", value)
			return
		}
		p.Port = uint16(value)
		return
	case string:
		// todo: extend to support other formats: iana-assigned port numbers, port ranges, etc.
		var digit uint64
		digit, err = strconv.ParseUint(value, 10, 16)
		if err != nil {
			err = fmt.Errorf("invalid port number: %v", value)
			return
		}
		p.Port = uint16(digit)
		return
	default:
		err = fmt.Errorf("invalid port number: %v", b)
		return
	}
}
