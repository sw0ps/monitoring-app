package domain

type ChannelList struct {
	list []ChannelItem
}

func (c *ChannelList) AddChannelItem(item ChannelItem) {
	c.list = append(c.list, item)
}

func (c *ChannelList) GetList() []ChannelItem {
	return c.list
}

//
//func (c ChannelList) AddChannelItem(item ChannelItem) {
//	c = append(c, item)
//}
//
//func (c ChannelList) GetList() []ChannelItem {
//	return c
//}