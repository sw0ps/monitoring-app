package service

import (
	"github.com/sw0ps/monitoring-app/internal/domain"
	"github.com/sw0ps/monitoring-app/internal/repository"
)

type Channel interface {
	CreateList(list domain.ChannelList) error
	GetAllList() (domain.ChannelList, error)
	UpdateList(list domain.ChannelList) error
}

type Service struct {
	Channel
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Channel: NewChannelService(repo.Channel),
	}
}
