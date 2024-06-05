package callbacks

import "encoding/json"

// 用户等级信息更新通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/vip/callback/member/usergrade_update.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcVipGradeInfoUpdate{})
}

type ChannelsEcVipGradeInfoUpdate struct {
	CreateTime   int  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	UserInfo     struct {
		ExperienceValue int `json:"experience_value"`
		Grade           int `json:"grade"`
	} `json:"user_info"`
}

func (ChannelsEcVipGradeInfoUpdate) GetMessageType() string {
	return "event"
}

func (ChannelsEcVipGradeInfoUpdate) GetEventType() string {
	return "channels_ec_vip_grade_info_update"
}

func (m ChannelsEcVipGradeInfoUpdate) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcVipGradeInfoUpdate) ParseFromJson(data []byte) (CallbackExtraInfoInterface,error) {
	var temp ChannelsEcVipGradeInfoUpdate
	err := json.Unmarshal(data, &temp)
	return temp, err
}
