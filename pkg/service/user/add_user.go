package userservice

import (
	"Dapp_2/pkg/database/models"
	"context"
)

func (userService *Service) Add(ctx context.Context, user *models.User) error {
	_, err := userService.DB.Insert(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
