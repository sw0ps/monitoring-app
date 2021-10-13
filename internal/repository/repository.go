package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sw0ps/monitoring-app/internal/domain"
)

type Channel interface {
	GetAllList() (domain.ChannelList, error)
	Create(channel domain.ChannelItem) (int, error)
	Update(channel domain.ChannelItem) (int, error)
}

type Repository struct {
	Channel
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Channel: NewChannelRepository(db),
	}
}