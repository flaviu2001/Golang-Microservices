package main

import (
	"Bleenco/common"
	"Bleenco/common/errors"
	"encoding/json"
	"fmt"
	"os"
)

func GetPorts() (entriesChannel chan common.Entry, errorChannel chan error) {
	entriesChannel = make(chan common.Entry, common.ChannelSize)
	errorChannel = make(chan error, common.ChannelSize)
	go func() {
		defer close(entriesChannel)
		defer close(errorChannel)
		file, err := os.Open(common.PortsJsonFilename)
		if err != nil {
			errorChannel <- &errors.PortError{Text: "Unable to open ports.json"}
			return
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "Error closing file")
			}
		}(file)
		decoder := json.NewDecoder(file)
		_, err = decoder.Token()
		if err != nil {
			errorChannel <- &errors.PortError{Text: "Unable to read starting token"}
			return
		}
		for decoder.More() {
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
