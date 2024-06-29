package repository

import (
	"antrein/dd-queue-auth/application/common/resource"
	"antrein/dd-queue-auth/internal/repository/config"
	"antrein/dd-queue-auth/internal/repository/room"
	cfg "antrein/dd-queue-auth/model/config"
)

type CommonRepository struct {
	ConfigRepo *config.Repository
	RoomRepo   *room.Repository
}

func NewCommonRepository(cfg *cfg.Config, rsc *resource.CommonResource) (*CommonRepository, error) {
	configRepo := config.New(cfg, rsc.Redis, rsc.GRPC)
	roomRepo := room.New(cfg, rsc.Redis)

	commonRepo := CommonRepository{
		ConfigRepo: configRepo,
		RoomRepo:   roomRepo,
	}
	return &commonRepo, nil
}
