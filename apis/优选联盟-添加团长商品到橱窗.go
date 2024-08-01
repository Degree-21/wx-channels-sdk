package apis

import (
	"encoding/json"
	
)

// 添加团长商品到橱窗
// 文档：https://developers.weixin.qq.com/doc/channels/API/leagueheadsupplier/addwindow.html

type ReqLeagueHeadsupplierWindowAdd struct {
	Appid        string `json:"appid"`
	FinderID     string `json:"finder_id"`
	Openfinderid string `json:"openfinderid"`
	ProductID    int  `json:"product_id"`
}



var _ bodyer = ReqLeagueHeadsupplierWindowAdd{}

func (x ReqLeagueHeadsupplierWindowAdd) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}


type RespLeagueHeadsupplierWindowAdd struct {
	CommonResp
}

var _ bodyer = RespLeagueHeadsupplierWindowAdd{}

func (x RespLeagueHeadsupplierWindowAdd) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecLeagueHeadsupplierWindowAdd(req ReqLeagueHeadsupplierWindowAdd) (RespLeagueHeadsupplierWindowAdd, error) {
	var resp RespLeagueHeadsupplierWindowAdd
	err := c.executeWXApiPost("/channels/ec/league/headsupplier/window/add", req, &resp, true)
	if err != nil {
		return RespLeagueHeadsupplierWindowAdd{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespLeagueHeadsupplierWindowAdd{}, bizErr
	}
	return resp, nil
}