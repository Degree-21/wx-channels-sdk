package callbacks

import "encoding/json"

// 用户积分更新通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/vip/callback/score/user_score_update.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcVipScoreUpdate{})
}

type ChannelsEcVipScoreUpdate struct {
	CreateTime   int  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	UserInfo     struct {
		DeltaScore int `json:"delta_score"`
		FlowType   int `json:"flow_type"`
		Score      int `json:"score"`
	} `json:"user_info"`
}

func (ChannelsEcVipScoreUpdate) GetMessageType() string {
	return "event"
}

func (ChannelsEcVipScoreUpdate) GetEventType() string {
	return "channels_ec_vip_score_update"
}

func (m ChannelsEcVipScoreUpdate) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcVipScoreUpdate) ParseFromJson(data []byte) (CallbackExtraInfoInterface,error) {
	var temp ChannelsEcVipScoreUpdate
	err := json.Unmarshal(data, &temp)
	return temp, err
}
