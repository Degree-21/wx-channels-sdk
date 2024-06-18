package apis

import (
	"encoding/json"
)

// 视频号获取当前的直播记录
// 文档：https://developers.weixin.qq.com/doc/channels/API/live/getfinderliverecordlist.html

type ReqFinderliveGetfinderliverecordlist struct{}

var _ bodyer = ReqFinderliveGetfinderliverecordlist{}

func (x ReqFinderliveGetfinderliverecordlist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespFinderliveGetfinderliverecordlist struct {
	CommonResp
	LiveList []struct {
		CoverURL    string `json:"cover_url"`
		Description string `json:"description"`
		ExportID    string `json:"export_id"`
		HeadURL     string `json:"head_url"`
		Nickname    string `json:"nickname"`
	} `json:"live_list"`
}

var _ bodyer = RespFinderliveGetfinderliverecordlist{}

func (x RespFinderliveGetfinderliverecordlist) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecFinderliveGetfinderliverecordlist(req ReqFinderliveGetfinderliverecordlist) (RespFinderliveGetfinderliverecordlist, error) {
	var resp RespFinderliveGetfinderliverecordlist
	err := c.executeWXApiPost("/channels/ec/finderlive/getfinderliverecordlist", req, &resp, true)
	if err != nil {
		return RespFinderliveGetfinderliverecordlist{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespFinderliveGetfinderliverecordlist{}, bizErr
	}
	return resp, nil
}
