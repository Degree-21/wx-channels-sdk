package apis

import (
	"encoding/json"
)

// 小店给分享员生成推广直播间的url
// 文档：https://developers.weixin.qq.com/doc/channels/API/sharer/promote_finder_live.html

type ReqSharerGetShopFinderLiveSharerUrl struct {
	ExportID string `json:"export_id"`
	FinderID string `json:"finder_id"`
}

var _ bodyer = ReqSharerGetShopFinderLiveSharerUrl{}

func (x ReqSharerGetShopFinderLiveSharerUrl) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespSharerGetShopFinderLiveSharerUrl struct {
	CommonResp
	URL string `json:"url"`
}

var _ bodyer = RespSharerGetShopFinderLiveSharerUrl{}

func (x RespSharerGetShopFinderLiveSharerUrl) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecSharerGetShopFinderLiveSharerUrl(req ReqSharerGetShopFinderLiveSharerUrl) (RespSharerGetShopFinderLiveSharerUrl, error) {
	var resp RespSharerGetShopFinderLiveSharerUrl
	err := c.executeWXApiPost("/channels/ec/sharer/get_shop_finder_live_sharer_url", req, &resp, true)
	if err != nil {
		return RespSharerGetShopFinderLiveSharerUrl{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSharerGetShopFinderLiveSharerUrl{}, bizErr
	}
	return resp, nil
}
