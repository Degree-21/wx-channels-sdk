package callbacks

import "encoding/json"

// 用户加入会员通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/vip/callback/member/user_join_vip.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcVipJoin{})
}

type ChannelsEcVipJoin struct {
	CreateTime   int  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	UserInfo     struct {
		JoinTime    int  `json:"join_time"`
		PhoneNumber string `json:"phone_number"`
	} `json:"user_info"`
}

func (ChannelsEcVipJoin) GetMessageType() string {
	return "event"
}

func (ChannelsEcVipJoin) GetEventType() string {
	return "channels_ec_vip_join"
}

func (m ChannelsEcVipJoin) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcVipJoin) ParseFromJson(data []byte) (CallbackExtraInfoInterface,error) {
	var temp ChannelsEcVipJoin
	err := json.Unmarshal(data, &temp)
	return temp, err
}
