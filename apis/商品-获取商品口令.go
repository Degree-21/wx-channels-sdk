package apis

import (
	"encoding/json"
	
)

// 获取商品口令
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/get_taglink.html

type ReqProductTaglinkGet struct {
	ProductID string `json:"product_id"`
}



var _ bodyer = ReqProductTaglinkGet{}

func (x ReqProductTaglinkGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}


type RespProductTaglinkGet struct {
	CommonResp
	ProductTaglink string `json:"product_taglink"`
}

var _ bodyer = RespProductTaglinkGet{}

func (x RespProductTaglinkGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecProductTaglinkGet(req ReqProductTaglinkGet) (RespProductTaglinkGet, error) {
	var resp RespProductTaglinkGet
	err := c.executeWXApiPost("/channels/ec/product/taglink/get", req, &resp, true)
	if err != nil {
		return RespProductTaglinkGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProductTaglinkGet{}, bizErr
	}
	return resp, nil
}