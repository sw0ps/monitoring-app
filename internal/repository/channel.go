package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sw0ps/monitoring-app/internal/domain"
)

const tableName = "channels_stats"

type ChannelRepository struct {
	db *sqlx.DB
}

func NewChannelRepository(db *sqlx.DB) *ChannelRepository {
	return &ChannelRepository{db: db}
}

func (r *ChannelRepository) GetAllList() (domain.ChannelList, error) {
	var list []domain.ChannelItem

	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	err := r.db.Select(&list, query)

	var outputList domain.ChannelList

	for _, item := range list {
		outputList.AddChannelItem(item)
	}

	return outputList, err
}

func (r *ChannelRepository) Create(channel domain.ChannelItem) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (channel_id, title, view_count, subscriber_count, video_count) VALUES ($1, $2, $3, $4, $5) RETURNING id", tableName)

	row := r.db.QueryRow(query, channel.ChannelId, channel.Title, channel.ViewCount, channel.SubscriberCount, channel.VideoCount)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ChannelRepository) Update(channel domain.ChannelItem) (int, error) {
	var outputId int

	query := fmt.Sprintf("UPDATE %s SET title = $1, view_count = $2, subscriber_count = $3, video_count = $4 WHERE channel_id = $5 RETURNING id", tableName)

	row := r.db.QueryRow(query, channel.Title, channel.ViewCount, channel.SubscriberCount, channel.VideoCount, channel.ChannelId)
	if err := row.Scan(&outputId); err != nil {
		return 0, err
	}

	return outputId, nil
}