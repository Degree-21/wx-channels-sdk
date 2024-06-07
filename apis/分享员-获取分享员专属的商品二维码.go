package apis

import (
	"encoding/json"
)

// 获取分享员专属的商品二维码
// 文档：https://developers.weixin.qq.com/doc/channels/API/sharer/get_sharer_product_qrcode.html

type ReqSharerGetSharerProductQrcode struct {
	Openid    string `json:"openid"`
	ProductID string `json:"product_id"`
}

var _ bodyer = ReqSharerGetSharerProductQrcode{}

func (x ReqSharerGetSharerProductQrcode) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespSharerGetSharerProductQrcode struct {
	CommonResp
	ProductQrcode string `json:"product_qrcode"`
}

var _ bodyer = RespSharerGetSharerProductQrcode{}

func (x RespSharerGetSharerProductQrcode) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecSharerGetSharerProductQrcode(req ReqSharerGetSharerProductQrcode) (RespSharerGetSharerProductQrcode, error) {
	var resp RespSharerGetSharerProductQrcode
	err := c.executeWXApiPost("/channels/ec/sharer/get_sharer_product_qrcode", req, &resp, true)
	if err != nil {
		return RespSharerGetSharerProductQrcode{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSharerGetSharerProductQrcode{}, bizErr
	}
	return resp, nil
}
