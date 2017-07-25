package gnet

import (
	guid "github.com/satori/go.uuid"
	exp "guacamole_client_go"
	"guacamole_client_go/gio"
)

// DelegatingGuacamoleTunnel ==> GuacamoleTunnel
//  * GuacamoleTunnel implementation which simply delegates all function calls to
//  * an underlying GuacamoleTunnel.
type DelegatingGuacamoleTunnel struct {
	tunnel GuacamoleTunnel
}

// NewDelegatingGuacamoleTunnel Construct
func NewDelegatingGuacamoleTunnel(tunnel GuacamoleTunnel) (ret DelegatingGuacamoleTunnel) {
	ret.tunnel = tunnel
	return
}

// GetUUID override GuacamoleTunnel.GetUUID
func (opt *DelegatingGuacamoleTunnel) GetUUID() guid.UUID {
	return opt.tunnel.GetUUID()
}

// GetSocket override GuacamoleTunnel.GetSocket
func (opt *DelegatingGuacamoleTunnel) GetSocket() GuacamoleSocket {
	return opt.tunnel.GetSocket()
}

// AcquireReader override GuacamoleTunnel.AcquireReader
func (opt *DelegatingGuacamoleTunnel) AcquireReader() gio.GuacamoleReader {
	return opt.tunnel.AcquireReader()
}

// ReleaseReader override GuacamoleTunnel.ReleaseReader
func (opt *DelegatingGuacamoleTunnel) ReleaseReader() {
	opt.tunnel.ReleaseReader()
}

// HasQueuedReaderThreads override GuacamoleTunnel.HasQueuedReaderThreads
func (opt *DelegatingGuacamoleTunnel) HasQueuedReaderThreads() bool {
	return opt.tunnel.HasQueuedReaderThreads()
}

// AcquireWriter override GuacamoleTunnel.AcquireWriter
func (opt *DelegatingGuacamoleTunnel) AcquireWriter() gio.GuacamoleWriter {
	return opt.tunnel.AcquireWriter()
}

// ReleaseWriter override GuacamoleTunnel.ReleaseWriter
func (opt *DelegatingGuacamoleTunnel) ReleaseWriter() {
	opt.tunnel.ReleaseWriter()
}

// HasQueuedWriterThreads override GuacamoleTunnel.HasQueuedWriterThreads
func (opt *DelegatingGuacamoleTunnel) HasQueuedWriterThreads() bool {
	return opt.tunnel.HasQueuedWriterThreads()
}

// Close override GuacamoleTunnel.Close
func (opt *DelegatingGuacamoleTunnel) Close() exp.ExceptionInterface {
	return opt.tunnel.Close()
}

// IsOpen override GuacamoleTunnel.IsOpen
func (opt *DelegatingGuacamoleTunnel) IsOpen() bool {
	if opt.tunnel != nil {
		return opt.tunnel.IsOpen()
	}
	return false
}
