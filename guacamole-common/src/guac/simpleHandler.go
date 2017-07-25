package main

import (
	"fmt"
	"guacamole_client_go/gservlet"
	"net/http"
)

// BaseToHTTPServletResponseInterface response
type BaseToHTTPServletResponseInterface struct {
	core     http.ResponseWriter
	commited bool
	err      error
}

// SwapResponse override
func SwapResponse(core http.ResponseWriter) gservlet.HTTPServletResponseInterface {
	return &BaseToHTTPServletResponseInterface{core: core}
}

// IsCommitted override
func (opt *BaseToHTTPServletResponseInterface) IsCommitted() (bool, error) {
	return opt.commited, opt.err
}

// AddHeader override
func (opt *BaseToHTTPServletResponseInterface) AddHeader(key, value string) {
	opt.core.Header().Add(key, value)
}

// SetHeader override
func (opt *BaseToHTTPServletResponseInterface) SetHeader(key, value string) {
	opt.core.Header().Set(key, value)
}

// SetContentType override
func (opt *BaseToHTTPServletResponseInterface) SetContentType(value string) {
	opt.core.Header().Set("Content-Type", value)
}

// SetContentLength override
func (opt *BaseToHTTPServletResponseInterface) SetContentLength(length int) {
	opt.core.Header().Set("Content-Length", fmt.Sprintf("%v", length))
}

// SendError override
func (opt *BaseToHTTPServletResponseInterface) SendError(sc int) error {
	opt.commited = true
	opt.core.WriteHeader(sc)
	return nil
}

// WriteString override
func (opt *BaseToHTTPServletResponseInterface) WriteString(data string) error {
	opt.commited = true
	opt.core.Write([]byte(data))
	return nil
}

// Write override
func (opt *BaseToHTTPServletResponseInterface) Write(data []byte) error {
	opt.commited = true
	opt.core.Write(data)
	return nil
}

// FlushBuffer override
func (opt *BaseToHTTPServletResponseInterface) FlushBuffer() error {
	if v, ok := opt.core.(http.Flusher); ok {
		v.Flush()
	}
	return nil
}

// Close override
func (opt *BaseToHTTPServletResponseInterface) Close() error {
	return opt.FlushBuffer()
}

// BaseToHTTPServletRequestInterface request
type BaseToHTTPServletRequestInterface struct {
	core *http.Request
}

// SwapRequest convert http.Request into Interface
func SwapRequest(core *http.Request) gservlet.HTTPServletRequestInterface {
	return &BaseToHTTPServletRequestInterface{core: core}
}

// GetQueryString override
func (opt *BaseToHTTPServletRequestInterface) GetQueryString() string {
	return opt.core.URL.RawQuery
}

// Read override
func (opt *BaseToHTTPServletRequestInterface) Read(buffer []byte) (int, error) {
	return opt.core.Body.Read(buffer)
}
