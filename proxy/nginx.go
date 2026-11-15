package main

type NginxServer struct {
	application     *Application
	maxAllowRequest int
	rateLimiter     map[string]int
}

func newNginxServer() *NginxServer {
	return &NginxServer{
		application:     &Application{},
		maxAllowRequest: 2,
		rateLimiter:     make(map[string]int),
	}
}

func (n *NginxServer) handleRequest(url, schema string) (int, string) {
	if !n.checkRateLimiting(url) {
		return 403, "Not Allowed"
	}

	return n.application.handleRequest(url, schema)
}

func (n *NginxServer) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] >= n.maxAllowRequest {
		return false
	}
	n.rateLimiter[url]++
	return true
}
