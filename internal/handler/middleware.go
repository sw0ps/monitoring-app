package handler

import (
	"github.com/sw0ps/monitoring-app/internal/domain"
	"github.com/sw0ps/monitoring-app/pkg/youtube"
	"log"
	"os"
)

// FirstRun add data after first run
func (h *Handler) FirstRun() error {
	list, err := h.services.GetAllList()

	if err != nil {
		return err
	}

	if len(list.GetList()) > 0 {
		return nil
	}

	youtubeList := tryYoutubeRequest()
	if err := h.services.CreateList(youtubeList); err != nil {
		return err
	}

	return nil
}

func tryYoutubeRequest() domain.ChannelList {

	var list domain.ChannelList

	items, err := youtube.TryRequest(youtube.Config{
		YoutubeKey: os.Getenv("YOUTUBE_KEY"),
		ChannelId:  os.Getenv("CHANNEL_ID"),
		Part:       os.Getenv("PART"),
	})
	if err != nil {
		log.Fatalf("%s", err)
	}

	for _, v := range items {
		list.AddChannelItem(
			domain.ChannelItem{
				ChannelId:       v.ChannelId,
				Title:           v.Snippet.Title,
				ViewCount:       v.Stats.ViewCount,
				SubscriberCount: v.Stats.SubscriberCount,
				VideoCount:      v.Stats.VideoCount,
			})
	}

	return list
}
