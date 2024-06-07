package apis

import (
	"encoding/json"
)

// 获取分享员专属的商品H5短链
// 文档：https://developers.weixin.qq.com/doc/channels/API/sharer/get_sharer_product_h5url.html

type ReqSharerGetSharerProductH5Url struct {
	Openid    string `json:"openid"`
	ProductID string `json:"product_id"`
}

var _ bodyer = ReqSharerGetSharerProductH5Url{}

func (x ReqSharerGetSharerProductH5Url) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespSharerGetSharerProductH5Url struct {
	CommonResp
	ProductH5url string `json:"product_h5url"`
}

var _ bodyer = RespSharerGetSharerProductH5Url{}

func (x RespSharerGetSharerProductH5Url) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecSharerGetSharerProductH5Url(req ReqSharerGetSharerProductH5Url) (RespSharerGetSharerProductH5Url, error) {
	var resp RespSharerGetSharerProductH5Url
	err := c.executeWXApiPost("/channels/ec/sharer/get_sharer_product_h5url", req, &resp, true)
	if err != nil {
		return RespSharerGetSharerProductH5Url{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespSharerGetSharerProductH5Url{}, bizErr
	}
	return resp, nil
}
