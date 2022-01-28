package service

import "Bleenco/port-domain-service/utils"

// Service This interface specifies the required operations that communicate with the repository
type Service interface {
	// Upsert This method attempts to insert the port and upon an unlocs conflict it will be updated
	Upsert(port utils.Port)

	// Select This method returns a number of pages corresponding to the specified page
	Select(page int) []utils.Port
}
