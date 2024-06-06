package apis

import (
	"encoding/json"
	
)

// 同意用户修改收货地址申请
// 文档：https://developers.weixin.qq.com/doc/channels/API/order/accept_address_modify_apply.html

type ReqOrderAddressmodifyAccept struct {
	OrderID string `json:"order_id"`
}



var _ bodyer = ReqOrderAddressmodifyAccept{}

func (x ReqOrderAddressmodifyAccept) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}


type RespOrderAddressmodifyAccept struct {
	CommonResp
}

var _ bodyer = RespOrderAddressmodifyAccept{}

func (x RespOrderAddressmodifyAccept) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecOrderAddressmodifyAccept(req ReqOrderAddressmodifyAccept) (RespOrderAddressmodifyAccept, error) {
	var resp RespOrderAddressmodifyAccept
	err := c.executeWXApiPost("/channels/ec/order/addressmodify/accept", req, &resp, true)
	if err != nil {
		return RespOrderAddressmodifyAccept{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespOrderAddressmodifyAccept{}, bizErr
	}
	return resp, nil
}