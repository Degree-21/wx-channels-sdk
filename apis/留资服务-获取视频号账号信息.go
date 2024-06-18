package apis

import (
	"encoding/json"
)

// 获取视频号账号信息
// 文档：https://developers.weixin.qq.com/doc/channels/API/live/get_finder_attr_by_appid.html

type ReqGetFinderAttrByAppid struct{}

var _ bodyer = ReqGetFinderAttrByAppid{}

func (x ReqGetFinderAttrByAppid) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetFinderAttrByAppid struct {
	CommonResp
	FinderAttr struct {
		FansCount int    `json:"fans_count"`
		Nickname  string `json:"nickname"`
		UniqID    string `json:"uniq_id"`
	} `json:"finder_attr"`
}

var _ bodyer = RespGetFinderAttrByAppid{}

func (x RespGetFinderAttrByAppid) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetFinderAttrByAppid(req ReqGetFinderAttrByAppid) (RespGetFinderAttrByAppid, error) {
	var resp RespGetFinderAttrByAppid
	err := c.executeWXApiPost("/channels/finderlive/get_finder_attr_by_appid", req, &resp, true)
	if err != nil {
		return RespGetFinderAttrByAppid{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetFinderAttrByAppid{}, bizErr
	}
	return resp, nil
}
