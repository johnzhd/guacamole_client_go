package gio

// Move FilteredGuacamoleReader from protocol folder to here
// Avoid cross depends

import (
	exp "guacamole_client_go"
	"guacamole_client_go/gprotocol"
)

// FilteredGuacamoleReader ==> GuacamoleReader
type FilteredGuacamoleReader struct {
	/**
	 * The wrapped GuacamoleReader.
	 */
	reader GuacamoleReader

	/**
	 * The filter to apply when reading instructions.
	 */
	filter gprotocol.GuacamoleFilter
}

/*NewFilteredGuacamoleReader *
* Wraps the given GuacamoleReader, applying the given filter to all read
* instructions. Future reads will return only instructions which pass
* the filter.
*
* @param reader The GuacamoleReader to wrap.
* @param filter The filter which dictates which instructions are read, and
*               how.
 */
func NewFilteredGuacamoleReader(
	reader GuacamoleReader,
	filter gprotocol.GuacamoleFilter,
) (ret FilteredGuacamoleReader) {
	ret.reader = reader
	ret.filter = filter
	return
}

// Available override GuacamoleReader.Available
func (opt *FilteredGuacamoleReader) Available() (ok bool, err exp.ExceptionInterface) {
	return opt.reader.Available()
}

// Read override GuacamoleReader.Read
func (opt *FilteredGuacamoleReader) Read() (ret []byte, err exp.ExceptionInterface) {
	filteredInstruction, err := opt.ReadInstruction()
	if err != nil {
		return
	}
	ret = []byte(filteredInstruction.String())
	return
}

// ReadInstruction override GuacamoleReader.ReadInstruction
func (opt *FilteredGuacamoleReader) ReadInstruction() (ret gprotocol.GuacamoleInstruction, err exp.ExceptionInterface) {
	var filteredInstruction gprotocol.GuacamoleInstruction

	for len(filteredInstruction.GetOpcode()) == 0 {
		filteredInstruction, err = opt.reader.ReadInstruction()
		if err != nil {
			return
		}
		if len(filteredInstruction.GetOpcode()) == 0 {
			return
		}
		filteredInstruction, err = opt.filter.Filter(filteredInstruction)
		if err != nil {
			return
		}
	}
	ret = filteredInstruction
	return
}
