package parser

import (
	"Bleenco/common"
	"Bleenco/common/errors"
	"encoding/json"
	"fmt"
	"os"
)

// GetPorts This function returns two channels, one for each entry parsed and one for encountered errors.
// It parses (with a buffered reader) the json specified in the filename parameter and feeds the entry channel
// with each entry found. Upon encountering one error the whole method halts and no further entries are fed.
// The method is non-blocking, and the channels receive their data in a background thread that produces said data.
func GetPorts(filename string) (entriesChannel chan common.Entry, errorChannel chan error) {
	entriesChannel = make(chan common.Entry, common.ChannelSize)
	errorChannel = make(chan error, common.ChannelSize)

	go func() {
		defer close(entriesChannel)
		defer close(errorChannel)

		file, err := os.Open(filename)
		if err != nil {
			errorChannel <- &errors.PortError{Text: fmt.Sprintf("Unable to open %s", filename)}
			return
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, fmt.Sprintf("Error closing %s", filename))
			}
		}(file)

		decoder := json.NewDecoder(file)

		// Skip over the first "{" token
		_, err = decoder.Token()

		if err != nil {
			errorChannel <- &errors.PortError{Text: "Unable to read starting token"}
			return
		}

		// While there are tokens left in the json
		for decoder.More() {
			// Reads the first field which is equal to the one in unlocs
			token, err := decoder.Token()

			if err != nil {
				errorChannel <- &errors.PortError{Text: "Unable to read port name"}
				return
			}

			entry := common.Entry{}

			switch token.(type) {
			case string:
				entry.PortName = token.(string)
			default:
				errorChannel <- &errors.PortError{Text: "Unexpected token type"}
				return
			}

			// Read the whole port from the json
			err = decoder.Decode(&entry.Port)

			if err != nil {
				errorChannel <- &errors.PortError{Text: "Unable to read port"}
				return
			}

			entriesChannel <- entry
		}
	}()
	return
}
