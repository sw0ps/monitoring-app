package domain

type ChannelItem struct {
	Id       		string `db:"id"`
	ChannelId       string `json:"id" db:"channel_id" binding:"required"`
	Title           string `json:"title" db:"title" binding:"required"`
	ViewCount       string `json:"viewCount" db:"view_count" binding:"required"`
	SubscriberCount string `json:"subscriberCount" db:"subscriber_count" binding:"required"`
	VideoCount      string `json:"videoCount" db:"video_count" binding:"required"`
}

func NewChannelItem(channelId, title, viewCount, subscriberCount, videoCount string) *ChannelItem {
	return &ChannelItem{
		Title:           title,
		ChannelId:       channelId,
		ViewCount:       viewCount,
		SubscriberCount: subscriberCount,
		VideoCount:      videoCount,
	}
}
