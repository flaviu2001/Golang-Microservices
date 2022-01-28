package service

import (
	"Bleenco/port-domain-service/constants"
	"Bleenco/port-domain-service/repository"
	"Bleenco/port-domain-service/utils"
)

type Impl struct {
	Repository repository.Repository
}

// Upsert Thid method removes the aliases and regions of a port to ensure that inserts will succeed after its upsertion.
func (i *Impl) Upsert(port utils.Port) {
	unlocs := port.Unlocs[0]

	i.Repository.RemoveAliases(unlocs)
	i.Repository.RemoveRegions(unlocs)

	var coord1 interface{}
	var coord2 interface{}
	if len(port.Coordinates) != 0 {
		coord1 = port.Coordinates[0]
		coord2 = port.Coordinates[1]
	} else {
		coord1 = nil
		coord2 = nil
	}

	portId := i.Repository.GetNewPortId()
	i.Repository.UpsertPort(portId, unlocs, port.Name, port.City, port.Country, coord1, coord2, port.Province, port.Timezone, port.Code)
	portId = i.Repository.FindPortId(unlocs)

	for _, alias := range port.Alias {
		i.Repository.InsertAlias(portId, unlocs, alias)
	}

	for _, region := range port.Regions {
		i.Repository.InsertRegion(portId, unlocs, region)
	}
}

func (i *Impl) Select(page int) []utils.Port {
	lowerBound := page * constants.PageSize
	upperBound := (page+1)*constants.PageSize - 1
	return i.Repository.SelectPorts(lowerBound, upperBound)
}
