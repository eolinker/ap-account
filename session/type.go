package session

import "time"

const (
	SessionName = "session"
	ExpireTime  = 72 * time.Hour
)

type Status int

const (
	NotLogin Status = iota
	Login
	Expired
)
