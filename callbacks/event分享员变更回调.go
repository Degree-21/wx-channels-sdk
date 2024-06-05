package callbacks

import "encoding/json"

// 分享员变更回调
// 文档: https://developers.weixin.qq.com/doc/channels/API/sharer/callback/channels_ec_sharer_change.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcSharerChange{})
}

type ChannelsEcSharerChange struct {
	CreateTime   int  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	BindStatus   int  `json:"bind_status"`
	Openid       string `json:"openid"`
	SharerType   int  `json:"sharer_type"`
}

func (ChannelsEcSharerChange) GetMessageType() string {
	return "event"
}

func (ChannelsEcSharerChange) GetEventType() string {
	return "channels_ec_sharer_change"
}

func (m ChannelsEcSharerChange) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcSharerChange) ParseFromJson(data []byte) (CallbackExtraInfoInterface,error) {
	var temp ChannelsEcSharerChange
	err := json.Unmarshal(data, &temp)
	return temp, err
}
