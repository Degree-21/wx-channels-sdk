# 生成API代码
# go run generate/api.go -doc=https://developers.weixin.qq.com/doc/channels/API/order/accept_address_modify_apply.html  -prefix=订单
api:
	go run generate/api.go -doc=$(doc) -prefix=$(prefix)

# 生成回调事件代码
# go run generate/callback.go -doc=https://developers.weixin.qq.com/doc/channels/API/sharer/callback/channels_ec_sharer_change.html  -prefix=event
callback:
	go run generate/callback.go -doc=$(doc)