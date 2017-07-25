package gio

import (
	exp "guacamole_client_go"
	"guacamole_client_go/gprotocol"
	"net"
)

// WriterGuacamoleWriter A GuacamoleWriter which wraps a standard Java Writer,
// using that Writer as the Guacamole instruction stream.
type WriterGuacamoleWriter struct {
	output *Stream
}

// NewWriterGuacamoleWriter Constuct function
//  * Creates a new WriterGuacamoleWriter which will use the given Writer as
//  * the Guacamole instruction stream.
//  *
//  * @param output The Writer to use as the Guacamole instruction stream.
func NewWriterGuacamoleWriter(output *Stream) (ret GuacamoleWriter) {
	one := WriterGuacamoleWriter{}
	one.output = output
	ret = &one
	return
}

// Write override GuacamoleWriter.Write
func (opt *WriterGuacamoleWriter) Write(chunk []byte, off, l int) (err exp.ExceptionInterface) {
	if len(chunk) < off+l {
		err = exp.GuacamoleServerException.Throw("Input buffer size smaller than required")
		return
	}
	e := opt.WriteAll(chunk[off : off+l])
	if e != nil {
		// Socket timeout will close so ...
		err = exp.GuacamoleConnectionClosedException.Throw("Connection to guacd is closed.", e.Error())
	}
	return
}

// WriteAll override GuacamoleWriter.WriteAll
func (opt *WriterGuacamoleWriter) WriteAll(chunk []byte) (err exp.ExceptionInterface) {
	_, e := opt.output.Write(chunk)
	if e == nil {
		return
	}
	switch e.(type) {
	case net.Error:
		ex := e.(net.Error)
		if ex.Timeout() {
			err = exp.GuacamoleUpstreamTimeoutException.Throw("Connection to guacd timed out.", e.Error())
		} else {
			err = exp.GuacamoleConnectionClosedException.Throw("Connection to guacd is closed.", e.Error())
		}
	default:
		err = exp.GuacamoleServerException.Throw(e.Error())
	}
	return
}

// WriteInstruction override GuacamoleWriter.WriteInstruction
func (opt *WriterGuacamoleWriter) WriteInstruction(instruction gprotocol.GuacamoleInstruction) (err exp.ExceptionInterface) {
	return opt.WriteAll([]byte(instruction.String()))
}
