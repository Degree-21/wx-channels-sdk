package apis

import (
	"encoding/json"
	"net/url"
)

// 获取店铺基本信息
// 文档：https://developers.weixin.qq.com/doc/channels/API/basics/getbasicinfo.html

type ReqBasicsInfoGet struct{}

var _ urlValuer = ReqBasicsInfoGet{}

func (x ReqBasicsInfoGet) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

type RespBasicsInfoGet struct {
	CommonResp
	Info struct {
		HeadimgURL  string `json:"headimg_url"`
		Nickname    string `json:"nickname"`
		Status      string `json:"status"`
		SubjectType string `json:"subject_type"`
		Username    string `json:"username"`
	} `json:"info"`
}

var _ bodyer = RespBasicsInfoGet{}

func (x RespBasicsInfoGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecBasicsInfoGet(req ReqBasicsInfoGet) (RespBasicsInfoGet, error) {
	var resp RespBasicsInfoGet
	err := c.executeWXApiGet("/channels/ec/basics/info/get", req, &resp, true)
	if err != nil {
		return RespBasicsInfoGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespBasicsInfoGet{}, bizErr
	}
	return resp, nil
}
