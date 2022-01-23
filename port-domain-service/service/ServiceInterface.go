package service

import "Bleenco/common"

type Service interface {
	Upsert(port common.Port)
	Select(page int) []common.Port
}
