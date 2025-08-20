package userservice

import (
	"Dapp_2/pkg/database"
	"Dapp_2/pkg/service"
)

type Service struct {
	*service.BaseService
}

func NewService(_db database.DB, _env string) *Service {
	return &Service{
		BaseService: &service.BaseService{
			DB:  _db,
			ENV: _env,
		},
	}
}
