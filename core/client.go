// core/client.go
package core

import (
	"net/http"
	"time"
)

var UserAgent = "Mozilla/5.0 (compatible; URLiesBot/1.0; +https://github.com/ryo/urlies)"

func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 60 * time.Second,
	}
}
