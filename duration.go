package marshallable

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
)

// Duration is a time.Duration wraper that marshals to and from a JSON string.
type Duration struct{ time.Duration }

// MarshalJSON implements the json.Marshaler interface.
func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Duration.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *Duration) UnmarshalJSON(b []byte) (err error) {
	var v interface{}
	if err = json.Unmarshal(b, &v); err != nil {
		return
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(math.Abs(value))
		return
	case string:
		d.Duration, err = time.ParseDuration(value) // todo: extend to support other formats: day and week
		return
	default:
		err = fmt.Errorf("invalid duration: %v", b)
		return
	}
}
