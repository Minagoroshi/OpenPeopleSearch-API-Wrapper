package openpeoplesearch

import "time"

type Session struct {
	Username    string
	Token       string
	TokenExpiry time.Time
}
