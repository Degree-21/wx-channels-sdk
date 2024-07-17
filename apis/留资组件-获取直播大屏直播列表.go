package apis

import (
	"encoding/json"
)

// 获取直播大屏直播列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/livedashboard/getlivelist.html

type ReqGetlivelist struct{}

var _ bodyer = ReqGetlivelist{}

func (x ReqGetlivelist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetlivelist struct {
	CommonResp
	Ds int `json:"ds"`
}

var _ bodyer = RespGetlivelist{}

func (x RespGetlivelist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetlivelist(req ReqGetlivelist) (RespGetlivelist, error) {
	var resp RespGetlivelist
	err := c.executeWXApiPost("/channels/livedashboard/getlivelist", req, &resp, true)
	if err != nil {
		return RespGetlivelist{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetlivelist{}, bizErr
	}
	return resp, nil
}
