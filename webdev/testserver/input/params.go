package input

import (
	"flag"
	"time"
)

type Params struct {
	verbose         bool
	statusCode      int
	responseHeaders map[string]string
	sleepTime       time.Duration
}

var paramsSingleton *Params

func (p *Params) loadParams() {
	flag.BoolVar(&p.verbose, "v", false, "Verbose mode. Usage: -v")
	flag.IntVar(&p.statusCode, "status_code", 200, "Http response status code. Usage: -status_code=X where X can be one of the valid status code (2xx, 4xx etc.)")

	var headers string
	flag.StringVar(&headers, "resp_headers", "", "Comma separated list of additional response headers. Usage: -resp_headers=\"key1:val1,key2:val2\"")

	var dur string
	flag.StringVar(&dur, "sleep", "", "Sleep duration. Usage: -sleep=5s. Currently only sec and millisec units are supported. Eg. 5s, 10ms")

	flag.Parse()

	p.responseHeaders = parseHeaders(headers)
	p.sleepTime = parseTimeStringIntoDuration(dur)
}

func (p *Params) Verbose() bool {
	return p.verbose
}

func (p *Params) StatusCode() int {
	return p.statusCode
}

func (p *Params) ResponseHeaders() map[string]string {
	return p.responseHeaders
}

func (p *Params) SleepTime() time.Duration {
	return p.sleepTime
}

func init() {
	paramsSingleton = &Params{}
	paramsSingleton.loadParams()
}

func GetParams() *Params {
	return paramsSingleton
}
