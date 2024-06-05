package callbacks

import "encoding/json"

// 用户积分兑换通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/vip/callback/score/user_score_exchange.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcVipScoreExchange{})
}

type ChannelsEcVipScoreExchange struct {
	CreateTime   int  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	ExchangeInfo struct {
		CouponInfo struct {
			RelatedCouponID int `json:"related_coupon_id"`
		} `json:"coupon_info"`
		PayScore    int `json:"pay_score"`
		ProductInfo struct {
			RelatedProductID int `json:"related_product_id"`
		} `json:"product_info"`
		ScoreItemType int `json:"score_item_type"`
	} `json:"exchange_info"`
}

func (ChannelsEcVipScoreExchange) GetMessageType() string {
	return "event"
}

func (ChannelsEcVipScoreExchange) GetEventType() string {
	return "channels_ec_vip_score_exchange"
}

func (m ChannelsEcVipScoreExchange) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcVipScoreExchange) ParseFromJson(data []byte) (CallbackExtraInfoInterface,error) {
	var temp ChannelsEcVipScoreExchange
	err := json.Unmarshal(data, &temp)
	return temp, err
}
