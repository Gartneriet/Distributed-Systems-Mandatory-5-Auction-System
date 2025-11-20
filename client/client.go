package client

import ()

type client struct {
	name                string
	clock               int32
	primaryServerPort   string
	secondaryServerPort string
}

func NewClient(name string) *client {
	return &client{name: name, clock: 0, primaryServerPort: "5000", secondaryServerPort: "6000"}
}
