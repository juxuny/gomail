package gomail

import (
	"github.com/gamexg/proxyclient"
)

var (
	proxyClient proxyclient.ProxyClient
)

// SetSocksProxy is a function that update proxy config
func SetProxy(proxyAddress string) error {
	var err error
	proxyClient, err = proxyclient.NewProxyClient(proxyAddress)
	if err != nil {
		return err
	}
	return nil
}
