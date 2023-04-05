package openpeoplesearch

import (
	"net/http"
	"time"
)

var (
	opsClient = &http.Client{Timeout: 10 * time.Second}
)
