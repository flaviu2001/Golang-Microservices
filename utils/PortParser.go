package utils

import (
	"Bleenco/errors"
	"encoding/json"
	"fmt"
	"os"
)

type Port struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float32 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

type Entry struct {
	PortName string
	Port     Port
}

func GetPorts() (entriesChannel chan Entry, errorChannel chan error) {
	entriesChannel = make(chan Entry, ChannelSize)
	errorChannel = make(chan error, ChannelSize)
	go func() {
		defer close(entriesChannel)
		defer close(errorChannel)
		file, err := os.Open(PortsJsonFilename)
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
			entry := Entry{}
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
