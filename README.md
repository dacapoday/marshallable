# Marshallable [![Go Reference](https://pkg.go.dev/badge/github.com/dacapoday/marshallable.svg)](https://pkg.go.dev/github.com/dacapoday/marshallable)
Make generic data types marshallable!

## Features
Implement methods:
* `String() (string)`
* `MarshalJSON() ([]byte, error)`
* `UnmarshalJSON([]byte) (error)`

for the following standard library types:
* `time.Duration`
* `net.IP`
* `url.URL`
* `mail.Address`
* `regexp.Regexp`

and common data types:
* TCP/UDP port number

## Install

```
go get github.com/dacapoday/marshallable
```

## Documentation

Read [Go Reference](https://pkg.go.dev/github.com/dacapoday/marshallable)

## Usage
Scenes: for configuration initialization and deserialization

```go
package main

import (
	"encoding/json"
	"fmt"

	m "github.com/dacapoday/marshallable"
)

type Config struct {
	ServerName    string        `json:"server_name"`
	ServerSecret  []byte        `json:"server_secret"`
	Timeout       m.Duration    `json:"timeout"`
	ListenPort    m.Port        `json:"listen_port"`
	ListenAddress m.IP          `json:"listen_address"`
	ApiServer     m.URL         `json:"api_server"`
	Contact       m.MailAddress `json:"contact"`
	SpamRule      m.Regexp      `json:"spam_rule"`
}

var base Config

func init() {
	// initialize Config with constant values
	base = Config{
		ServerName:    "server-name",
		ServerSecret:  []byte("binary encoded with base64"),
		Timeout:       m.Duration{10 * time.Second},
		ListenPort:    m.Port{8080},
		ListenAddress: m.MustIP("10.0.0.1"),
		ApiServer:     m.MustURL("http://api.example.com"),
		Contact:       m.MustMailAddress("who@example.com"),
		SpamRule:      m.MustRegexp("^[a-zA-Z0-9]{4,}$"),
	}
}

func main() {
	fmt.Println("marshal:")
	data, err := json.Marshal(base)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)

	fmt.Println("unmarshal:")
	config := new(Config)
	err = json.Unmarshal(data, config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", config)
}

```

## License

[MIT License](LICENSE)
