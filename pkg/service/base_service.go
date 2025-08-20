package service

import "Dapp_2/pkg/database"

type BaseService struct {
	DB  database.DB
	ENV string
}
