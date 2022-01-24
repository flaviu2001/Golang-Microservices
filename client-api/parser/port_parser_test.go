package parser

import (
	"Bleenco/common"
	"testing"
)

func TestGetPorts(t *testing.T) {
	entriesChannel, errorChannel := GetPorts(common.PortsJsonFilenameTest)
	entriesOpen := true
	errorOpen := true
	running := true
	var err error
	var entry common.Entry
	var entries = make([]common.Entry, 0)
	for running {
		select {
		case entry, entriesOpen = <-entriesChannel:
			if entriesOpen {
				entries = append(entries, entry)
			} else {
				entriesChannel = nil
			}
		case err, errorOpen = <-errorChannel:
			if errorOpen {
				t.Fatalf("Unexpected error: %s", err)
			} else {
				errorChannel = nil
			}
		default:
			if entriesChannel == nil && errorChannel == nil {
				running = false
			}
		}
	}
	if len(entries) != 1 {
		t.Fatalf("Entries length should be 1")
	}
	port := entries[0].Port
	if !(port.Name == "Ajman" && port.City == "Ajman" && port.Country == "United Arab Emirates" &&
		port.Alias[0] == "alias1" && len(port.Regions) == 0 && len(port.Coordinates) == 2 && port.Province == "Ajman" &&
		port.Timezone == "Asia/Dubai" && port.Unlocs[0] == "AEAJM" && port.Code == "52000") {
		t.Fatalf("Unexpected port")
	}
}
