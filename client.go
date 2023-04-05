package openpeoplesearch

import (
	"net/http"
	"time"
)

var (
	opsClient = &http.Client{Timeout: 10 * time.Second}
)

const (
	USERAGENT = "OPSClientGO/1.0"
)

type Session struct {
	Username    string
	Token       string
	TokenExpiry time.Time
}
