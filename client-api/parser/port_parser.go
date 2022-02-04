package parser

import (
	"Bleenco/client-api/constants"
	"Bleenco/client-api/errors"
	"Bleenco/client-api/utils"
	"encoding/json"
	"fmt"
	"os"
)

// GetPorts This function returns two channels, one for each entry parsed and one for encountered errors.
// It parses (with a buffered reader) the json specified in the filename parameter and feeds the entry channel
// with each entry found. Upon encountering one error the whole method halts and no further entries are fed.
// The method is non-blocking, and the channels receive their data in a background thread that produces said data.
// The end of parsing is marked by a nil value in the error channel.
func GetPorts(filename string) (entriesChannel chan utils.Entry, errorChannel chan error) {
	entriesChannel = make(chan utils.Entry, constants.ChannelSize)
	errorChannel = make(chan error, constants.ChannelSize)

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

			entry := utils.Entry{}

			switch token.(type) {
			case string:
				entry.PortName = token.(string)
			default:
				errorChannel <- &errors.PortError{Text: "Unexpected token type"}
				return
			}

			// Read the whole port from the json
			if err = decoder.Decode(&entry.Port); err != nil {
				errorChannel <- &errors.PortError{Text: "Unable to read port"}
				return
			}

			entriesChannel <- entry
		}
	}()
	return
}
