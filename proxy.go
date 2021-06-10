package gomail

import (
	"fmt"
	"github.com/thinkgos/go-socks5/ccsocks5"
	"golang.org/x/net/proxy"
)

var (
	socks5Client *ccsocks5.Client
)

// SetSocksProxy is a function that update proxy config
func SetSocks5Proxy(ip string, port int, auth ...proxy.Auth) {
	var opts = make([]ccsocks5.Option, 0)
	if len(auth) > 0 {
		opts = append(opts, ccsocks5.WithAuth(&auth[0]))
	}
	socks5Client = ccsocks5.NewClient(fmt.Sprintf("%s:%d", ip, port), opts...)
}
