package apis

import (
	"encoding/json"
)

// 获取订单列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/order/list_get.html

type ReqOrderListGet struct {
	CreateTimeRange *struct {
		EndTime   int `json:"end_time,omitempty"`
		StartTime int `json:"start_time,omitempty"`
	} `json:"create_time_range,omitempty"`
	UpdateTimeRange *struct {
		EndTime   int `json:"end_time,omitempty"`
		StartTime int `json:"start_time,omitempty"`
	} `json:"update_time_range,omitempty"`
	NextKey  string `json:"next_key"`
	PageSize int    `json:"page_size"`
}

var _ bodyer = ReqOrderListGet{}

func (x ReqOrderListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespOrderListGet struct {
	CommonResp
	HasMore     bool     `json:"has_more"`
	NextKey     string   `json:"next_key"`
	OrderIDList []string `json:"order_id_list"`
}

var _ bodyer = RespOrderListGet{}

func (x RespOrderListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecOrderListGet(req ReqOrderListGet) (RespOrderListGet, error) {
	var resp RespOrderListGet
	err := c.executeWXApiPost("/channels/ec/order/list/get", req, &resp, true)
	if err != nil {
		return RespOrderListGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespOrderListGet{}, bizErr
	}
	return resp, nil
}
