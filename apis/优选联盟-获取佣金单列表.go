package apis

import (
	"encoding/json"
	
)

// 获取佣金单列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/leagueheadsupplier/getorderlist.html

type ReqLeagueHeadsupplierOrderListGet struct {
	CreateTimeRange struct {
		EndTime   int `json:"end_time"`
		StartTime int `json:"start_time"`
	} `json:"create_time_range"`
	PageSize int `json:"page_size"`
}



var _ bodyer = ReqLeagueHeadsupplierOrderListGet{}

func (x ReqLeagueHeadsupplierOrderListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}


type RespLeagueHeadsupplierOrderListGet struct {
	CommonResp
	HasMore bool   `json:"has_more"`
	List    []struct {
		OrderID string `json:"order_id"`
		SkuID   string `json:"sku_id"`
	} `json:"list"`
}

var _ bodyer = RespLeagueHeadsupplierOrderListGet{}

func (x RespLeagueHeadsupplierOrderListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecLeagueHeadsupplierOrderListGet(req ReqLeagueHeadsupplierOrderListGet) (RespLeagueHeadsupplierOrderListGet, error) {
	var resp RespLeagueHeadsupplierOrderListGet
	err := c.executeWXApiPost("/channels/ec/league/headsupplier/order/list/get", req, &resp, true)
	if err != nil {
		return RespLeagueHeadsupplierOrderListGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespLeagueHeadsupplierOrderListGet{}, bizErr
	}
	return resp, nil
}