package apis

import (
	"encoding/json"
)

// 获取分享员订单列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/sharer/get_sharer_order_list.html

type ReqSharerGetSharerOrderList struct {
	EndTime    int    `json:"end_time"`
	Openid     string `json:"openid"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
	ShareScene int    `json:"share_scene"`
	StartTime  int    `json:"start_time"`
}

var _ bodyer = ReqSharerGetSharerOrderList{}

func (x ReqSharerGetSharerOrderList) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespSharerGetSharerOrderList struct {
	CommonResp
	OrderList []struct {
		FinderSceneInfo struct {
			FinderNickname string `json:"finder_nickname"`
			LiveExportID   string `json:"live_export_id"`
			PromoterID     string `json:"promoter_id"`
			VideoExportID  string `json:"video_export_id"`
			VideoTitle     string `json:"video_title"`
		} `json:"finder_scene_info"`
		FromWecom    bool   `json:"from_wecom"`
		OrderID      string `json:"order_id"`
		ProductID    string `json:"product_id"`
		ShareScene   int    `json:"share_scene"`
		SharerOpenid string `json:"sharer_openid"`
		SharerType   int    `json:"sharer_type"`
		SkuID        string `json:"sku_id"`
	} `json:"order_list"`
}

var _ bodyer = RespSharerGetSharerOrderList{}

func (x RespSharerGetSharerOrderList) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecSharerGetSharerOrderList(req ReqSharerGetSharerOrderList) (RespSharerGetSharerOrderList, error) {
	var resp RespSharerGetSharerOrderList
	err := c.executeWXApiPost("/channels/ec/sharer/get_sharer_order_list", req, &resp, true)
	if err != nil {
		return RespSharerGetSharerOrderList{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSharerGetSharerOrderList{}, bizErr
	}
	return resp, nil
}
