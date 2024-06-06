package apis

import (
	"encoding/json"
	
)

// 修改物流信息
// 文档：https://developers.weixin.qq.com/doc/channels/API/order/deliveryinfo_update.html

type ReqOrderDeliveryinfoUpdate struct {
	DeliveryList []struct {
		DeliverType  int  `json:"deliver_type"`
		DeliveryID   string `json:"delivery_id"`
		ProductInfos []struct {
			ProductCnt int  `json:"product_cnt"`
			ProductID  string `json:"product_id"`
			SkuID      string `json:"sku_id"`
		} `json:"product_infos"`
		WaybillID string `json:"waybill_id"`
	} `json:"delivery_list"`
	OrderID string `json:"order_id"`
}



var _ bodyer = ReqOrderDeliveryinfoUpdate{}

func (x ReqOrderDeliveryinfoUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}


type RespOrderDeliveryinfoUpdate struct {
	CommonResp
}

var _ bodyer = RespOrderDeliveryinfoUpdate{}

func (x RespOrderDeliveryinfoUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecOrderDeliveryinfoUpdate(req ReqOrderDeliveryinfoUpdate) (RespOrderDeliveryinfoUpdate, error) {
	var resp RespOrderDeliveryinfoUpdate
	err := c.executeWXApiPost("/channels/ec/order/deliveryinfo/update", req, &resp, true)
	if err != nil {
		return RespOrderDeliveryinfoUpdate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespOrderDeliveryinfoUpdate{}, bizErr
	}
	return resp, nil
}