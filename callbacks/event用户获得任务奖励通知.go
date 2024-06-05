package callbacks

import "encoding/json"

// 用户获得任务奖励通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/vip/callback/member/user_finish_task.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcVipTaskAward{})
}

type ChannelsEcVipTaskAward struct {
	CreateTime   int  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	TaskInfo     struct {
		AwardInfo struct {
			AwardType int `json:"award_type"`
			CouponID  int `json:"coupon_id"`
			GainScore int `json:"gain_score"`
		} `json:"award_info"`
		TaskID   int  `json:"task_id"`
		TaskName string `json:"task_name"`
	} `json:"task_info"`
}

func (ChannelsEcVipTaskAward) GetMessageType() string {
	return "event"
}

func (ChannelsEcVipTaskAward) GetEventType() string {
	return "channels_ec_vip_task_award"
}

func (m ChannelsEcVipTaskAward) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcVipTaskAward) ParseFromJson(data []byte) (CallbackExtraInfoInterface,error) {
	var temp ChannelsEcVipTaskAward
	err := json.Unmarshal(data, &temp)
	return temp, err
}
