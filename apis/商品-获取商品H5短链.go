package apis

import (
	"encoding/json"
	
)

// 获取商品H5短链
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/get_h5url.html

type ReqProductH5UrlGet struct {
	ProductID string `json:"product_id"`
}



var _ bodyer = ReqProductH5UrlGet{}

func (x ReqProductH5UrlGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}


type RespProductH5UrlGet struct {
	CommonResp
	ProductH5url string `json:"product_h5url"`
}

var _ bodyer = RespProductH5UrlGet{}

func (x RespProductH5UrlGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecProductH5UrlGet(req ReqProductH5UrlGet) (RespProductH5UrlGet, error) {
	var resp RespProductH5UrlGet
	err := c.executeWXApiPost("/channels/ec/product/h5url/get", req, &resp, true)
	if err != nil {
		return RespProductH5UrlGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProductH5UrlGet{}, bizErr
	}
	return resp, nil
}