package callbacks

import "encoding/json"

// 用户注销会员通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/vip/callback/member/user_close_vip.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcVipClose{})
}

type ChannelsEcVipClose struct {
	CreateTime   int  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	UserInfo     struct {
		CloseTime int `json:"close_time"`
	} `json:"user_info"`
}

func (ChannelsEcVipClose) GetMessageType() string {
	return "event"
}

func (ChannelsEcVipClose) GetEventType() string {
	return "channels_ec_vip_close"
}

func (m ChannelsEcVipClose) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcVipClose) ParseFromJson(data []byte) (CallbackExtraInfoInterface,error) {
	var temp ChannelsEcVipClose
	err := json.Unmarshal(data, &temp)
	return temp, err
}
