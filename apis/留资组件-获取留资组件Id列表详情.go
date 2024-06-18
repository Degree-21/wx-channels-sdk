package apis

import (
	"encoding/json"
)

// 获取留资组件Id列表详情
// 文档：https://developers.weixin.qq.com/doc/channels/API/leads/get_leads_component_id.html

type ReqGetLeadsComponentId struct {
	LastBuffer string `json:"last_buffer"`
}

var _ bodyer = ReqGetLeadsComponentId{}

func (x ReqGetLeadsComponentId) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetLeadsComponentId struct {
	ContinueFlag bool `json:"continue_flag"`
	CommonResp
	Item []struct {
		LeadsComponentID string `json:"leads_component_id"`
		LeadsDescription string `json:"leads_description"`
		Status           int    `json:"status"`
	} `json:"item"`
	LastBuffer string `json:"last_buffer"`
}

var _ bodyer = RespGetLeadsComponentId{}

func (x RespGetLeadsComponentId) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetLeadsComponentId(req ReqGetLeadsComponentId) (RespGetLeadsComponentId, error) {
	var resp RespGetLeadsComponentId
	err := c.executeWXApiPost("/channels/leads/get_leads_component_id", req, &resp, true)
	if err != nil {
		return RespGetLeadsComponentId{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetLeadsComponentId{}, bizErr
	}
	return resp, nil
}
