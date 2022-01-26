package marshallable

import (
	"encoding/json"
	"errors"
	"net"
)

var ErrInvalidIP = errors.New("invalid IP")

type IP struct {
	net.IP
}

// String returns the IP as a string.
func (ip IP) String() string {
	if ip.IP == nil {
		return "0.0.0.0"
	}
	return ip.IP.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (ip IP) MarshalJSON() ([]byte, error) {
	return json.Marshal(ip.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ip *IP) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
	if err != nil {
		return
	}
	ip.IP = net.ParseIP(value)
	if ip.IP == nil {
		err = ErrInvalidIP
	}
	return
}

// MustIP returns the IP or panics.
func MustIP(value string) IP {
	ip := net.ParseIP(value)
	if ip == nil {
		panic(ErrInvalidIP)
	}
	return IP{ip}
}

var ErrInvalidIPv4 = errors.New("invalid IPv4")

type IPv4 struct {
	net.IP
}

// String returns the IP as a string.
func (ip IPv4) String() string {
	if ip.IP == nil {
		return "0.0.0.0"
	}
	return ip.IP.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (ip IPv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(ip.IP.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ip *IPv4) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
	if err != nil {
		return
	}
	ip.IP = net.ParseIP(value)
	if ip.IP.To4() == nil {
		err = ErrInvalidIPv4
	}
	return
}

// MustIPv4 returns the IPv4 or panics.
func MustIPv4(value string) IPv4 {
	ip := net.ParseIP(value)
	if ip.To4() == nil {
		panic(ErrInvalidIP)
	}
	return IPv4{ip}
}

var ErrInvalidIPv6 = errors.New("invalid IPv6")

type IPv6 struct {
	net.IP
}

// String returns the IP as a string.
func (ip IPv6) String() string {
	if ip.IP == nil {
		return "0:0:0:0:0:0:0:0"
	}
	return ip.IP.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (ip IPv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(ip.IP.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ip *IPv6) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
	if err != nil {
		return
	}
	ip.IP = net.ParseIP(value)
	if ip.IP.To16() == nil {
		err = ErrInvalidIPv6
	}
	return
}

// MustIPv6 returns the IPv6 or panics.
func MustIPv6(value string) IPv6 {
	ip := net.ParseIP(value)
	if ip.To16() == nil {
		panic(ErrInvalidIP)
	}
	return IPv6{ip}
}
