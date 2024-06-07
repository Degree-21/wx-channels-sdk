package apis

import (
	"encoding/json"
)

// 小店给分享员生成推广直播间预告的url
// 文档：https://developers.weixin.qq.com/doc/channels/API/sharer/promote_finder_notice.html

type ReqSharerGetShopFinderLiveNoticeSharerUrl struct {
	FinderID string `json:"finder_id"`
	NoticeID string `json:"notice_id"`
}

var _ bodyer = ReqSharerGetShopFinderLiveNoticeSharerUrl{}

func (x ReqSharerGetShopFinderLiveNoticeSharerUrl) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespSharerGetShopFinderLiveNoticeSharerUrl struct {
	CommonResp
	URL string `json:"url"`
}

var _ bodyer = RespSharerGetShopFinderLiveNoticeSharerUrl{}

func (x RespSharerGetShopFinderLiveNoticeSharerUrl) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecSharerGetShopFinderLiveNoticeSharerUrl(req ReqSharerGetShopFinderLiveNoticeSharerUrl) (RespSharerGetShopFinderLiveNoticeSharerUrl, error) {
	var resp RespSharerGetShopFinderLiveNoticeSharerUrl
	err := c.executeWXApiPost("/channels/ec/sharer/get_shop_finder_live_notice_sharer_url", req, &resp, true)
	if err != nil {
		return RespSharerGetShopFinderLiveNoticeSharerUrl{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSharerGetShopFinderLiveNoticeSharerUrl{}, bizErr
	}
	return resp, nil
}
