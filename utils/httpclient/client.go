package httpclient

import (
	"net/http"
	"time"
)

// Default is the default HTTP client to use.
var Default = &http.Client{
	Timeout: time.Second * 30,
}
