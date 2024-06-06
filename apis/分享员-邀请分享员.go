package apis

import (
	"encoding/json"
)

// 邀请分享员
// 文档：https://developers.weixin.qq.com/doc/channels/API/sharer/bindsharer.html

type ReqSharerBind struct {
	Username string `json:"username"`
}

var _ bodyer = ReqSharerBind{}

func (x ReqSharerBind) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespSharerBind struct {
	CommonResp
	QrcodeImg       string `json:"qrcode_img"`
	QrcodeImgBase64 string `json:"qrcode_img_base64"`
}

var _ bodyer = RespSharerBind{}

func (x RespSharerBind) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecSharerBind(req ReqSharerBind) (RespSharerBind, error) {
	var resp RespSharerBind
	err := c.executeWXApiPost("/channels/ec/sharer/bind", req, &resp, true)
	if err != nil {
		return RespSharerBind{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSharerBind{}, bizErr
	}
	return resp, nil
}
