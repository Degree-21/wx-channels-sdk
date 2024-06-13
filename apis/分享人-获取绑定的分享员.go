package apis

import (
	"encoding/json"
)

// 获取绑定的分享员
// 文档：https://developers.weixin.qq.com/doc/channels/API/sharer/search_sharer.html

type ReqSharerSearchSharer struct {
	Openid   string `json:"openid"`
	Username string `json:"username"`
}

var _ bodyer = ReqSharerSearchSharer{}

func (x ReqSharerSearchSharer) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespSharerSearchSharer struct {
	BindTime int `json:"bind_time"`
	CommonResp
	Nickname   string `json:"nickname"`
	Openid     string `json:"openid"`
	SharerType int    `json:"sharer_type"`
	Unionid    string `json:"unionid"`
}

var _ bodyer = RespSharerSearchSharer{}

func (x RespSharerSearchSharer) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecSharerSearchSharer(req ReqSharerSearchSharer) (RespSharerSearchSharer, error) {
	var resp RespSharerSearchSharer
	err := c.executeWXApiPost("/channels/ec/sharer/search_sharer", req, &resp, true)
	if err != nil {
		return RespSharerSearchSharer{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSharerSearchSharer{}, bizErr
	}
	return resp, nil
}
