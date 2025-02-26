package proxy

import (
	"fmt"
	"net/http/httputil"
	"net/url"
)

type ProxyReverse struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	Protocol     string `json:"protocol"`
	reverseProxy *httputil.ReverseProxy
	error_       error
}

// Create a new ProxyReverse parameter is a struct of type ProxyReverse with configured values for host, port and protocol
var NewProxyReverse = func(c ProxyReverse) *ProxyReverse {
	var host string = c.Host
	if c.Port != "" {
		host = fmt.Sprintf("%s:%s", c.Host, c.Port)
	}
	targetHost := fmt.Sprintf("%s://%s", c.Protocol, host)

	p := &ProxyReverse{
		Host:     c.Host,
		Port:     c.Port,
		Protocol: c.Protocol,
	}

	target, e := url.Parse(targetHost)
	if e != nil {
		p.error_ = errors.Errorf("Error parsing target host: %s Error: %s", targetHost, e.Error())
		return p
	}
	p.reverseProxy = httputil.NewSingleHostReverseProxy(target)
	return p
}

// RequestProxy is a function that returns a gin.HandlerFunc. Process the request of the router and pass it
// to the ReverseProxy of type httputil.ReverseProxy, which processes the request and returns the response to router
func (p *ProxyReverse) RequestProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		p.reverseProxy.ServeHTTP(c.Writer, c.Request)
	}
}
