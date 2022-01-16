package main

import (
	"Bleenco/utils"
	"fmt"
	"os"
)

func main() {
	entriesChannel, errorChannel := utils.GetPorts()
	entriesOpen := true
	errorOpen := true
	running := true
	var entry utils.Entry
	var err error
	for running {
		select {
		case entry, entriesOpen = <-entriesChannel:
			if entriesOpen {
				fmt.Println(entry)
			} else {
				entriesChannel = nil
			}
		case err, errorOpen = <-errorChannel:
			if errorOpen {
				_, _ = fmt.Fprintf(os.Stderr, "%s", err.Error())
			} else {
				errorChannel = nil
			}
		default:
			if entriesChannel == nil && errorChannel == nil {
				running = false
			}
		}
	}
}
