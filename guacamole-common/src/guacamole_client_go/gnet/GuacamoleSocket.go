package gnet

import (
	exp "guacamole_client_go"
	"guacamole_client_go/gio"
)

// GuacamoleSocket Provides abstract socket-like access to a Guacamole connection.
type GuacamoleSocket interface {
	/**
	 * Returns a GuacamoleReader which can be used to read from the
	 * Guacamole instruction stream associated with the connection
	 * represented by this GuacamoleSocket.
	 *
	 * @return A GuacamoleReader which can be used to read from the
	 *         Guacamole instruction stream.
	 */
	GetReader() gio.GuacamoleReader

	/**
	 * Returns a GuacamoleWriter which can be used to write to the
	 * Guacamole instruction stream associated with the connection
	 * represented by this GuacamoleSocket.
	 *
	 * @return A GuacamoleWriter which can be used to write to the
	 *         Guacamole instruction stream.
	 */
	GetWriter() gio.GuacamoleWriter

	/**
	 * Releases all resources in use by the connection represented by this
	 * GuacamoleSocket.
	 *
	 * @throws GuacamoleException If an error occurs while releasing resources.
	 */
	Close() exp.ExceptionInterface

	/**
	 * Returns whether this GuacamoleSocket is open and can be used for reading
	 * and writing.
	 *
	 * @return true if this GuacamoleSocket is open, false otherwise.
	 */
	IsOpen() bool
}
