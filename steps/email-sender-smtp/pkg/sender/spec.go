package sender

import (
	"net/mail"
)

type AddressSpec struct {
	*mail.Address
}

func (as *AddressSpec) UnmarshalText(data []byte) (err error) {
	as.Address, err = mail.ParseAddress(string(data))
	return
}

type ServerSpec struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	TLS      *bool  `json:"tls"`
}

type BodySpec struct {
	Text string `json:"text"`
	HTML string `json:"html"`
}

type Spec struct {
	Server         ServerSpec    `json:"server"`
	From           AddressSpec   `json:"from"`
	To             []AddressSpec `json:"to"`
	Cc             []AddressSpec `json:"cc"`
	Bcc            []AddressSpec `json:"bcc"`
	Subject        string        `json:"subject"`
	Body           BodySpec      `json:"body"`
	TimeoutSeconds uint          `json:"timeoutSeconds"`
}
