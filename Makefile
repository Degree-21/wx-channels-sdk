# 生成API代码
# go run generate/api.go -doc=https://developers.weixin.qq.com/doc/channels/API/basics/getbasicinfo.html   -prefix=基础
api:
	go run generate/api.go -doc=$(doc) -prefix=$(prefix)

# 生成回调事件代码
callback:
	go run generate/callback.go -doc=$(doc)