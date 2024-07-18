package main

import "time"

// message represents a single message in our chat
type message struct {
	Name      string
	Message   string
	When      time.Time
	AvatarURL string
}
