package callbacks

import "encoding/json"

// 小店注销回调事件
// 文档: https://developers.weixin.qq.com/doc/channels/API/basics/callback/channels_ec_basic_close_store.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcCloseStore{})
}

type ChannelsEcCloseStore struct {
	CreateTime     int  `json:"CreateTime"`
	Event          string `json:"Event"`
	FromUserName   string `json:"FromUserName"`
	MsgType        string `json:"MsgType"`
	ToUserName     string `json:"ToUserName"`
	BindAppid      string `json:"bind_appid"`
	CloseTimestamp string `json:"close_timestamp"`
}

func (ChannelsEcCloseStore) GetMessageType() string {
	return "event"
}

func (ChannelsEcCloseStore) GetEventType() string {
	return "channels_ec_close_store"
}

func (m ChannelsEcCloseStore) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcCloseStore) ParseFromJson(data []byte) (CallbackExtraInfoInterface,error) {
	var temp ChannelsEcCloseStore
	err := json.Unmarshal(data, &temp)
	return temp, err
}
