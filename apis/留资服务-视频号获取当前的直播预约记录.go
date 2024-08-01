package apis

import (
	"encoding/json"
)

// 视频号获取当前的直播预约记录
// 文档：https://developers.weixin.qq.com/doc/channels/API/live/getfinderlivenoticerecordlist.html

type ReqFinderliveGetfinderlivenoticerecordlist struct{}

var _ bodyer = ReqFinderliveGetfinderlivenoticerecordlist{}

func (x ReqFinderliveGetfinderlivenoticerecordlist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespFinderliveGetfinderlivenoticerecordlist struct {
	CommonResp
	LiveNoticeList []struct {
		Description   string `json:"description"`
		HeadURL       string `json:"head_url"`
		Nickname      string `json:"nickname"`
		NoticeID      string `json:"notice_id"`
		StartTime     int    `json:"start_time"`
		Status        int    `json:"status"`
		StatusWording string `json:"status_wording"`
	} `json:"live_notice_list"`
}

var _ bodyer = RespFinderliveGetfinderlivenoticerecordlist{}

func (x RespFinderliveGetfinderlivenoticerecordlist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecFinderliveGetfinderlivenoticerecordlist(req ReqFinderliveGetfinderlivenoticerecordlist) (RespFinderliveGetfinderlivenoticerecordlist, error) {
	var resp RespFinderliveGetfinderlivenoticerecordlist
	err := c.executeWXApiPost("/channels/ec/finderlive/getfinderlivenoticerecordlist", req, &resp, true)
	if err != nil {
		return RespFinderliveGetfinderlivenoticerecordlist{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespFinderliveGetfinderlivenoticerecordlist{}, bizErr
	}
	return resp, nil
}
