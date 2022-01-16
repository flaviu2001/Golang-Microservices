package errors

import "fmt"

type PortError struct {
	Text string
}

func (e *PortError) Error() string {
	return fmt.Sprintf("PortError: %v", e.Text)
}
