package service

import (
	"Bleenco/common"
	"Bleenco/port-domain-service/repository"
)

type Impl struct {
	Repository repository.Repository
}

// Upsert Thid method removes the aliases and regions of a port to ensure that inserts will succeed after its upsertion.
func (i *Impl) Upsert(port common.Port) {
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

func (i *Impl) Select(page int) []common.Port {
	lowerBound := page * common.PageSize
	upperBound := (page+1)*common.PageSize - 1
	return i.Repository.SelectPorts(lowerBound, upperBound)
}
