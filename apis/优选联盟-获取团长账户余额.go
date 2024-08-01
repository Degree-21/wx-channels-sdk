package apis

import (
	"encoding/json"
	
)

// 获取团长账户余额
// 文档：https://developers.weixin.qq.com/doc/channels/API/leagueheadsupplier/getbalance.html

type ReqLeagueHeadsupplierFundsBalanceGet struct{}



var _ bodyer = ReqLeagueHeadsupplierFundsBalanceGet{}

func (x ReqLeagueHeadsupplierFundsBalanceGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}


type RespLeagueHeadsupplierFundsBalanceGet struct {
	AvailableAmount int  `json:"available_amount"`
	CommonResp
	PendingAmount   int  `json:"pending_amount"`
}

var _ bodyer = RespLeagueHeadsupplierFundsBalanceGet{}

func (x RespLeagueHeadsupplierFundsBalanceGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecLeagueHeadsupplierFundsBalanceGet(req ReqLeagueHeadsupplierFundsBalanceGet) (RespLeagueHeadsupplierFundsBalanceGet, error) {
	var resp RespLeagueHeadsupplierFundsBalanceGet
	err := c.executeWXApiPost("/channels/ec/league/headsupplier/funds/balance/get", req, &resp, true)
	if err != nil {
		return RespLeagueHeadsupplierFundsBalanceGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespLeagueHeadsupplierFundsBalanceGet{}, bizErr
	}
	return resp, nil
}