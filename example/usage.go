package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

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
var empty Config

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

	fmt.Println("marshal empty value:")

	empty_data, err := json.Marshal(empty)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", empty_data)

	// use underlying type
	var fqdn *url.URL
	fqdn = config.ApiServer.URL

	var timeout time.Duration
	timeout = config.Timeout.Duration

	_ = config.ListenPort.Port
	_ = config.ListenAddress.IP
	_ = config.Contact.Address
	_ = config.SpamRule.Regexp

	_ = fqdn
	_ = timeout
}
