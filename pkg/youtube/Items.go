package youtube

type Items struct {
	ChannelId string  `json:"id"`
	Snippet   Snippet `json:"snippet"`
	Stats     Stats   `json:"statistics"`
}
