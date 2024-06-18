package apis

import (
	"encoding/json"
)

// 获取账号收集的留资数量
// 文档：https://developers.weixin.qq.com/doc/channels/API/live/get_finder_live_leads_data.html

type ReqGetFinderLiveLeadsData struct {
	EndTime    int `json:"end_time"`
	SourceType int `json:"source_type"`
	StartTime  int `json:"start_time"`
}

var _ bodyer = ReqGetFinderLiveLeadsData{}

func (x ReqGetFinderLiveLeadsData) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetFinderLiveLeadsData struct {
	CommonResp
	Item []struct {
		ComponentType int `json:"component_type"`
		LeadsCount    int `json:"leads_count"`
		TrafficType   int `json:"traffic_type"`
	} `json:"item"`
}

var _ bodyer = RespGetFinderLiveLeadsData{}

func (x RespGetFinderLiveLeadsData) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetFinderLiveLeadsData(req ReqGetFinderLiveLeadsData) (RespGetFinderLiveLeadsData, error) {
	var resp RespGetFinderLiveLeadsData
	err := c.executeWXApiPost("/channels/finderlive/get_finder_live_leads_data", req, &resp, true)
	if err != nil {
		return RespGetFinderLiveLeadsData{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetFinderLiveLeadsData{}, bizErr
	}
	return resp, nil
}
