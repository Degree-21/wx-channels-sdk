package apis

import (
	"encoding/json"
)

// 解绑分享员
// 文档：https://developers.weixin.qq.com/doc/channels/API/sharer/unbindsharer.html

type ReqSharerUnbind struct {
	OpenidList []string `json:"openid_list"`
}

var _ bodyer = ReqSharerUnbind{}

func (x ReqSharerUnbind) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespSharerUnbind struct {
	CommonResp
	FailOpenid    []interface{} `json:"fail_openid"`
	RefuseOpenid  []interface{} `json:"refuse_openid"`
	SuccessOpenid []interface{} `json:"success_openid"`
}

var _ bodyer = RespSharerUnbind{}

func (x RespSharerUnbind) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecSharerUnbind(req ReqSharerUnbind) (RespSharerUnbind, error) {
	var resp RespSharerUnbind
	err := c.executeWXApiPost("/channels/ec/sharer/unbind", req, &resp, true)
	if err != nil {
		return RespSharerUnbind{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSharerUnbind{}, bizErr
	}
	return resp, nil
}
