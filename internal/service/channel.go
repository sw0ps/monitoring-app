package service

import (
	"github.com/sw0ps/monitoring-app/internal/domain"
	"github.com/sw0ps/monitoring-app/internal/repository"
)

type ChannelService struct {
	repo repository.Channel
}

func NewChannelService(repo repository.Channel) *ChannelService {
	return &ChannelService{repo: repo}
}

func (s *ChannelService) CreateList(list domain.ChannelList) error {

	for _, item := range list.GetList() {
		if _, err := s.repo.Create(item); err != nil {
			return err
		}
	}

	return nil
}

func (s *ChannelService) GetAllList() (domain.ChannelList, error) {

	return s.repo.GetAllList()
}

func (s *ChannelService) UpdateList(list domain.ChannelList) error {

	for _, item := range list.GetList() {
		if _, err := s.repo.Update(item); err != nil {
			return err
		}
	}

	return nil
}
