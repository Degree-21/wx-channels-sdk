package apis

import (
	"encoding/json"
)

// 获取直播大屏数据
// 文档：https://developers.weixin.qq.com/doc/channels/API/livedashboard/getlivedata.html

type ReqGetlivedata struct {
	ExportID string `json:"export_id"`
}

var _ bodyer = ReqGetlivedata{}

func (x ReqGetlivedata) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespGetlivedata struct {
	CommonResp
	AverageWatchSecondsPerAudience int `json:"average_watch_seconds_per_audience"`
	CommentUv                      int `json:"comment_uv"`
	HotQuota                       int `json:"hot_quota"`
	ImpressionUv                   int `json:"impression_uv"`
	MaxOnlineWatchUv               int `json:"max_online_watch_uv"`
	NewFansClubUv                  int `json:"new_fans_club_uv"`
	NewFollowUv                    int `json:"new_follow_uv"`
	NewWatchUv                     int `json:"new_watch_uv"`
	RewardUv                       int `json:"reward_uv"`
	SharingUv                      int `json:"sharing_uv"`
}

var _ bodyer = RespGetlivedata{}

func (x RespGetlivedata) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecGetlivedata(req ReqGetlivedata) (RespGetlivedata, error) {
	var resp RespGetlivedata
	err := c.executeWXApiPost("/channels/livedashboard/getlivedata", req, &resp, true)
	if err != nil {
		return RespGetlivedata{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetlivedata{}, bizErr
	}
	return resp, nil
}
