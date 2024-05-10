package service

import (
	"github.com/DanielVieirass/um_help/config"
	"github.com/DanielVieirass/um_help/repo"
	"github.com/rs/zerolog"
)

type Service struct {
	User *UserService
}

func New(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *Service {
	return &Service{
		User: newUserService(cfg, logger, repo),
	}
}
