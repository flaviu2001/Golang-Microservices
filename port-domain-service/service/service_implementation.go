package service

import (
	"Bleenco/common"
	"Bleenco/port-domain-service/constants"
	"Bleenco/port-domain-service/repository"
)

type Impl struct {
	Repository repository.Repository
}

// Upsert This method removes the aliases and regions of a port to ensure that inserts will succeed after its upsertion.
func (i *Impl) Upsert(port common.Port) {
	var coord1 interface{}
	var coord2 interface{}
	if len(port.Coordinates) != 0 {
		coord1 = port.Coordinates[0]
		coord2 = port.Coordinates[1]
	} else {
		coord1 = nil
		coord2 = nil
	}

	i.Repository.UpsertPort(port.Unlocs[0], port.Name, port.City, port.Country, port.Alias, port.Regions, coord1, coord2, port.Province, port.Timezone, port.Code)
}

func (i *Impl) Select(page int) []common.Port {
	lowerBound := page * constants.PageSize
	upperBound := (page+1)*constants.PageSize - 1
	return i.Repository.SelectPorts(lowerBound, upperBound)
}
