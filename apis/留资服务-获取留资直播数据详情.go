package apis

import (
	"encoding/json"
)

// 获取留资直播数据详情
// 文档：https://developers.weixin.qq.com/doc/channels/API/live/get_finder_live_data_list.html

type ReqGetFinderLiveDataList struct {
	EndTime    int    `json:"end_time"`
	LastBuffer string `json:"last_buffer"`
	StartTime  int    `json:"start_time"`
}

var _ bodyer = ReqGetFinderLiveDataList{}

func (x ReqGetFinderLiveDataList) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetFinderLiveDataList struct {
	ContinueFlag bool `json:"continue_flag"`
	CommonResp
	Item []struct {
		AudiencesAvgSeconds   int    `json:"audiences_avg_seconds"`
		ExportID              string `json:"export_id"`
		ForwardCount          int    `json:"forward_count"`
		LiveDurationInSeconds int    `json:"live_duration_in_seconds"`
		LiveStartTime         int    `json:"live_start_time"`
		MaxOnlineCount        int    `json:"max_online_count"`
		NewFollowCount        int    `json:"new_follow_count"`
		NewFollowCountBiz     int    `json:"new_follow_count_biz"`
		TotalAudienceCount    int    `json:"total_audience_count"`
		TotalCheerCount       int    `json:"total_cheer_count"`
		TotalCommentCount     int    `json:"total_comment_count"`
	} `json:"item"`
	LastBuffer string `json:"last_buffer"`
}

var _ bodyer = RespGetFinderLiveDataList{}

func (x RespGetFinderLiveDataList) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetFinderLiveDataList(req ReqGetFinderLiveDataList) (RespGetFinderLiveDataList, error) {
	var resp RespGetFinderLiveDataList
	err := c.executeWXApiPost("/channels/finderlive/get_finder_live_data_list", req, &resp, true)
	if err != nil {
		return RespGetFinderLiveDataList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetFinderLiveDataList{}, bizErr
	}
	return resp, nil
}
