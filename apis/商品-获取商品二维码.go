package apis

import (
	"encoding/json"
	
)

// 获取商品二维码
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/get_qrcode.html

type ReqProductQrcodeGet struct {
	ProductID string `json:"product_id"`
}



var _ bodyer = ReqProductQrcodeGet{}

func (x ReqProductQrcodeGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}


type RespProductQrcodeGet struct {
	CommonResp
	ProductQrcode string `json:"product_qrcode"`
}

var _ bodyer = RespProductQrcodeGet{}

func (x RespProductQrcodeGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecProductQrcodeGet(req ReqProductQrcodeGet) (RespProductQrcodeGet, error) {
	var resp RespProductQrcodeGet
	err := c.executeWXApiPost("/channels/ec/product/qrcode/get", req, &resp, true)
	if err != nil {
		return RespProductQrcodeGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProductQrcodeGet{}, bizErr
	}
	return resp, nil
}