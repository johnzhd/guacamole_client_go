package gio

import (
	exp "guacamole_client_go"
	"guacamole_client_go/gprotocol"
)

// GuacamoleWriter Provides abstract and raw character write access
// to a stream of Guacamole instructions.
type GuacamoleWriter interface {
	// WriteEx function
	//  * Writes a portion of the given array of characters to the Guacamole
	//  * instruction stream. The portion must contain only complete Guacamole
	//  * instructions.
	//  *
	//  * @param chunk An array of characters containing Guacamole instructions.
	//  * @param off The start offset of the portion of the array to write.
	//  * @param len The length of the portion of the array to write.
	//  * @throws GuacamoleException If an error occurred while writing the
	//  *                            portion of the array specified.
	Write(chunk []byte, off, len int) (err exp.ExceptionInterface)

	// WriteAll
	//  * Writes the entire given array of characters to the Guacamole instruction
	//  * stream. The array must consist only of complete Guacamole instructions.
	//  *
	//  * @param chunk An array of characters consisting only of complete
	//  *              Guacamole instructions.
	//  * @throws GuacamoleException If an error occurred while writing the
	//  *                            the specified array.
	WriteAll(chunk []byte) (err exp.ExceptionInterface)

	// WriteInstruction function
	//  * Writes the given fully parsed instruction to the Guacamole instruction
	//  * stream.
	//  *
	//  * @param instruction The Guacamole instruction to write.
	//  * @throws GuacamoleException If an error occurred while writing the
	//  *                            instruction.
	WriteInstruction(instruction gprotocol.GuacamoleInstruction) (err exp.ExceptionInterface)
}
