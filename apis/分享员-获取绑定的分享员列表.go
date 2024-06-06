package apis

import (
	"encoding/json"
)

// 获取绑定的分享员列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/sharer/get_sharer_list.html

type ReqSharerGetSharerList struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	SharerType int `json:"sharer_type"`
}

var _ bodyer = ReqSharerGetSharerList{}

func (x ReqSharerGetSharerList) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespSharerGetSharerList struct {
	CommonResp
	SharerInfoList []struct {
		BindTime   int    `json:"bind_time"`
		Nickname   string `json:"nickname"`
		Openid     string `json:"openid"`
		SharerType int    `json:"sharer_type"`
		Unionid    string `json:"unionid"`
	} `json:"sharer_info_list"`
}

var _ bodyer = RespSharerGetSharerList{}

func (x RespSharerGetSharerList) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecSharerGetSharerList(req ReqSharerGetSharerList) (RespSharerGetSharerList, error) {
	var resp RespSharerGetSharerList
	err := c.executeWXApiPost("/channels/ec/sharer/get_sharer_list", req, &resp, true)
	if err != nil {
		return RespSharerGetSharerList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSharerGetSharerList{}, bizErr
	}
	return resp, nil
}
