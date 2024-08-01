package apis

import (
	"encoding/json"
	
)

// 获取合作小店列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/leagueheadsupplier/getshoplist.html

type ReqLeagueHeadsupplierShopListGet struct {
	PageSize int `json:"page_size"`
}



var _ bodyer = ReqLeagueHeadsupplierShopListGet{}

func (x ReqLeagueHeadsupplierShopListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}


type RespLeagueHeadsupplierShopListGet struct {
	CommonResp
	ShopList []struct {
		BaseInfo struct {
			Appid      string `json:"appid"`
			HeadimgURL string `json:"headimg_url"`
			Nickname   string `json:"nickname"`
		} `json:"base_info"`
		Status int `json:"status"`
	} `json:"shop_list"`
}

var _ bodyer = RespLeagueHeadsupplierShopListGet{}

func (x RespLeagueHeadsupplierShopListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecLeagueHeadsupplierShopListGet(req ReqLeagueHeadsupplierShopListGet) (RespLeagueHeadsupplierShopListGet, error) {
	var resp RespLeagueHeadsupplierShopListGet
	err := c.executeWXApiPost("/channels/ec/league/headsupplier/shop/list/get", req, &resp, true)
	if err != nil {
		return RespLeagueHeadsupplierShopListGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespLeagueHeadsupplierShopListGet{}, bizErr
	}
	return resp, nil
}