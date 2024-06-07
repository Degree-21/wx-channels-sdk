package apis

import (
	"encoding/json"
)

// 获取分享员专属的商品口令
// 文档：https://developers.weixin.qq.com/doc/channels/API/sharer/get_sharer_product_taglink.html

type ReqSharerGetSharerProductTaglink struct {
	Openid    string `json:"openid"`
	ProductID string `json:"product_id"`
}

var _ bodyer = ReqSharerGetSharerProductTaglink{}

func (x ReqSharerGetSharerProductTaglink) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespSharerGetSharerProductTaglink struct {
	CommonResp
	ProductTaglink string `json:"product_taglink"`
}

var _ bodyer = RespSharerGetSharerProductTaglink{}

func (x RespSharerGetSharerProductTaglink) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecSharerGetSharerProductTaglink(req ReqSharerGetSharerProductTaglink) (RespSharerGetSharerProductTaglink, error) {
	var resp RespSharerGetSharerProductTaglink
	err := c.executeWXApiPost("/channels/ec/sharer/get_sharer_product_taglink", req, &resp, true)
	if err != nil {
		return RespSharerGetSharerProductTaglink{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSharerGetSharerProductTaglink{}, bizErr
	}
	return resp, nil
}
