package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type proxy struct {
	liveUrlStr    string
	testUrlStr    string
	parsedTestUrl *url.URL
	parsedLiveUrl *url.URL
}

func NewProxy(liveUrlStr, testUrlStr string) proxy {
	parsedLiveUrl, _ := url.Parse(liveUrlStr)
	parsedTestUrl, _ := url.Parse(testUrlStr)

	inst := proxy{
		liveUrlStr:    liveUrlStr,
		testUrlStr:    testUrlStr,
		parsedLiveUrl: parsedLiveUrl,
		parsedTestUrl: parsedTestUrl,
	}
	return inst
}

func (p proxy) LiveUrl() string {
	return p.liveUrlStr
}

func (p proxy) TestUrl() string {
	return p.testUrlStr
}

func (p proxy) StartProxy(port string) {
	log.Printf( "Proxy starting on port %s", port)

	// start server
	http.HandleFunc("/", p.handleRequestAndRedirect)
	if err := http.ListenAndServe( ":" + port, nil); err != nil {
		panic(err)
	}
}

// Given a request send it to the appropriate url
func (p proxy) handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	logRequest(req)

	// decision point and comparison here ?
	p.serveReverseProxy(res, req)
}

// Serve a reverse proxy for a given url
func (p proxy) serveReverseProxy(res http.ResponseWriter, req *http.Request) {
	// parse the url
	url := p.parsedLiveUrl

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

// Log the request information
func logRequest(req *http.Request) {
	log.Printf("proxy_url: %s\n", req.RequestURI)
}
