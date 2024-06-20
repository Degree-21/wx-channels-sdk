package apis

import (
	"encoding/json"
)

// 获取订单详情
// 文档：https://developers.weixin.qq.com/doc/channels/API/order/get.html

type ReqOrderGet struct {
	OrderID             string `json:"order_id"`
	EncodeSensitiveInfo bool   `json:"encode_sensitive_info"`
}

var _ bodyer = ReqOrderGet{}

func (x ReqOrderGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespOrderGet struct {
	CommonResp
	Order struct {
		AftersaleDetail struct {
			AftersaleOrderList []struct {
				AftersaleOrderID string `json:"aftersale_order_id"`
				Status           int    `json:"status"`
			} `json:"aftersale_order_list"`
			OnAftersaleOrderCnt int `json:"on_aftersale_order_cnt"`
		} `json:"aftersale_detail"`
		CreateTime  int    `json:"create_time"`
		Openid      string `json:"openid"`
		OrderDetail struct {
			CouponInfo struct {
				UserCouponID string `json:"user_coupon_id"`
			} `json:"coupon_info"`
			DeliveryInfo struct {
				AddressInfo struct {
					CityName     string `json:"city_name"`
					CountyName   string `json:"county_name"`
					DetailInfo   string `json:"detail_info"`
					PostalCode   string `json:"postal_code"`
					ProvinceName string `json:"province_name"`
					TelNumber    string `json:"tel_number"`
					UserName     string `json:"user_name"`
				} `json:"address_info"`
				DeliverMethod       int `json:"deliver_method"`
				DeliveryProductInfo []struct {
					DeliverType  int    `json:"deliver_type"`
					DeliveryID   string `json:"delivery_id"`
					DeliveryTime int    `json:"delivery_time"`
					ProductInfos []struct {
						ProductCnt int    `json:"product_cnt"`
						ProductID  string `json:"product_id"`
						SkuID      string `json:"sku_id"`
					} `json:"product_infos"`
					WaybillID string `json:"waybill_id"`
				} `json:"delivery_product_info"`
				ShipDoneTime int `json:"ship_done_time"`
			} `json:"delivery_info"`
			ExtInfo struct {
				CustomerNotes string `json:"customer_notes"`
				FinderID      string `json:"finder_id"`
				LiveID        string `json:"live_id"`
				MerchantNotes string `json:"merchant_notes"`
				OrderScene    int    `json:"order_scene"`
			} `json:"ext_info"`
			PayInfo struct {
				PayTime       int    `json:"pay_time"`
				PaymentMethod int    `json:"payment_method"`
				PrepayID      string `json:"prepay_id"`
				PrepayTime    int    `json:"prepay_time"`
				TransactionID string `json:"transaction_id"`
			} `json:"pay_info"`
			PriceInfo struct {
				DiscountedPrice int  `json:"discounted_price"`
				Freight         int  `json:"freight"`
				IsDiscounted    bool `json:"is_discounted"`
				OrderPrice      int  `json:"order_price"`
				ProductPrice    int  `json:"product_price"`
			} `json:"price_info"`
			ProductInfos []struct {
				FinishAftersaleSkuCnt int `json:"finish_aftersale_sku_cnt"`
				MarketPrice           int `json:"market_price"`
				OnAftersaleSkuCnt     int `json:"on_aftersale_sku_cnt"`
				ProductID             int `json:"product_id"`
				SalePrice             int `json:"sale_price"`
				SkuAttrs              []struct {
					AttrKey   string `json:"attr_key"`
					AttrValue string `json:"attr_value"`
				} `json:"sku_attrs"`
				SkuCnt   int    `json:"sku_cnt"`
				SkuID    int    `json:"sku_id"`
				ThumbImg string `json:"thumb_img"`
				Title    string `json:"title"`
			} `json:"product_infos"`
			SettleInfo struct {
				CommissionFee        int `json:"commission_fee"`
				PredictCommissionFee int `json:"predict_commission_fee"`
				SettleTime           int `json:"settle_time"`
			} `json:"settle_info"`
			SharerInfo struct {
				HandlingProgress int    `json:"handling_progress"`
				ShareScene       int    `json:"share_scene"`
				SharerOpenid     string `json:"sharer_openid"`
				SharerType       int    `json:"sharer_type"`
				SharerUnionid    string `json:"sharer_unionid"`
			} `json:"sharer_info"`
			SkuSharerInfos []struct {
				ShareScene    int    `json:"share_scene"`
				SharerOpenid  string `json:"sharer_openid"`
				SharerType    int    `json:"sharer_type"`
				SharerUnionid string `json:"sharer_unionid"`
				SkuID         string `json:"sku_id"`
			} `json:"sku_sharer_infos"`
		} `json:"order_detail"`
		OrderID    string `json:"order_id"`
		Status     int    `json:"status"`
		UpdateTime int    `json:"update_time"`
	} `json:"order"`
}

var _ bodyer = RespOrderGet{}

func (x RespOrderGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecOrderGet(req ReqOrderGet) (RespOrderGet, error) {
	var resp RespOrderGet
	err := c.executeWXApiPost("/channels/ec/order/get", req, &resp, true)
	if err != nil {
		return RespOrderGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespOrderGet{}, bizErr
	}
	return resp, nil
}
