package main

import (
	"guacamole_client_go/gnet"
	"guacamole_client_go/gprotocol"
	"guacamole_client_go/gservlet"
	"net/http"
)

const enterPath = "/tunnel"

// ServletHandleStruct servlet
type ServletHandleStruct struct {
	gservlet.GuacamoleHTTPTunnelServlet
}

// NewServletHandleStruct servlet
func NewServletHandleStruct(doConnect gservlet.DoConnectInterface) (ret ServletHandleStruct) {
	ret.GuacamoleHTTPTunnelServlet = gservlet.NewGuacamoleHTTPTunnelServlet(doConnect)
	return
}

// GetEnterPath for http enter
func (opt *ServletHandleStruct) GetEnterPath() string {
	return enterPath
}

// ServeHTTP override
func (opt *ServletHandleStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != enterPath || r.Method != "GET" && r.Method != "POST" {
		http.NotFound(w, r)
		return
	}
	var err error
	response := SwapResponse(w)
	request := SwapRequest(r)

	response.SetHeader("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case "GET":
		err = opt.DoGet(request, response)
	case "POST":
		err = opt.DoPost(request, response)
	}
	if err != nil {
		http.NotFound(w, r)
		return
	}
	return
}

// DemoDoConnect Demo & test code
func DemoDoConnect(request gservlet.HTTPServletRequestInterface,
) (ret gnet.GuacamoleTunnel, err error) {
	config := gprotocol.NewGuacamoleConfiguration()
	infomation := gprotocol.NewGuacamoleClientInformation()

	// guac
	config.SetProtocol("vnc")
	config.SetParameter("color-depth", "8")
	config.SetParameter("cursor", "remote")
	// vm
	config.SetParameter("hostname", "127.0.0.1")
	config.SetParameter("port", "5902")

	// view
	infomation.SetOptimalScreenHeight(600)
	infomation.SetOptimalScreenWidth(800)

	core, err := gnet.NewInetGuacamoleSocket("127.0.0.1", 4822)
	if err != nil {
		return
	}
	socket, err := gnet.NewConfiguredGuacamoleSocket3(
		&core,
		config,
		infomation,
	)
	if err != nil {
		return
	}
	ret = gnet.NewSimpleGuacamoleTunnel(&socket)
	return
}
