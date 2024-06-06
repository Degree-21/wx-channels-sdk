package apis

type EnumCode int64

// OrderStatus + 待付款
const EnumCodeOrderStatus10 EnumCode = 10
		
// OrderStatus + 待发货
const EnumCodeOrderStatus20 EnumCode = 20
		
// OrderStatus + 部分发货
const EnumCodeOrderStatus21 EnumCode = 21
		
// OrderStatus + 待收货
const EnumCodeOrderStatus30 EnumCode = 30
		
// OrderStatus + 完成
const EnumCodeOrderStatus100 EnumCode = 100
		
// OrderStatus + 全部商品售后之后，订单取消
const EnumCodeOrderStatus200 EnumCode = 200
		
// OrderStatus + 未付款用户主动取消或超时未付款订单自动取消
const EnumCodeOrderStatus250 EnumCode = 250
		
// DeliveryType + 自寄快递
const EnumCodeDeliveryType1 EnumCode = 1
		
// DeliveryType + 在线签约快递单
const EnumCodeDeliveryType2 EnumCode = 2
		
// DeliveryType + 虚拟商品无需物流发货
const EnumCodeDeliveryType3 EnumCode = 3
		
// DeliveryType + 在线快递散单
const EnumCodeDeliveryType4 EnumCode = 4
		
// ShareScene + 直播间
const EnumCodeShareScene1 EnumCode = 1
		
// ShareScene + 橱窗
const EnumCodeShareScene2 EnumCode = 2
		
// ShareScene + 短视频
const EnumCodeShareScene3 EnumCode = 3
		
// ShareScene + 视频号主页
const EnumCodeShareScene4 EnumCode = 4
		
// ShareScene + 商品详情页
const EnumCodeShareScene5 EnumCode = 5
		
// ShareScene + 带商品的公众号文章
const EnumCodeShareScene6 EnumCode = 6
		
// ShareScene + 商品链接
const EnumCodeShareScene7 EnumCode = 7
		
// ShareScene + 商品二维码
const EnumCodeShareScene8 EnumCode = 8
		
// ShareScene + 商品短链
const EnumCodeShareScene9 EnumCode = 9
		
// ShareScene + 分享直播间
const EnumCodeShareScene10 EnumCode = 10
		
// ShareScene + 分享预约直播间
const EnumCodeShareScene11 EnumCode = 11
		
// ShareScene + 视频号橱窗的短链
const EnumCodeShareScene12 EnumCode = 12
		
// ShareScene + 视频号橱窗的二维码
const EnumCodeShareScene13 EnumCode = 13
		
// PaymentMethod + 微信支付
const EnumCodePaymentMethod1 EnumCode = 1
		
// PaymentMethod + 先用后付
const EnumCodePaymentMethod2 EnumCode = 2
		
// PaymentMethod + 抽奖商品0元订单
const EnumCodePaymentMethod3 EnumCode = 3
		
// PaymentMethod + 会员积分兑换订单
const EnumCodePaymentMethod4 EnumCode = 4
		
// InspectStatus + 待录入送检信息
const EnumCodeInspectStatus0 EnumCode = 0
		
// InspectStatus + 待送检
const EnumCodeInspectStatus1 EnumCode = 1
		
// InspectStatus + 未入库已取消
const EnumCodeInspectStatus2 EnumCode = 2
		
// InspectStatus + 入库异常
const EnumCodeInspectStatus3 EnumCode = 3
		
// InspectStatus + 已入库
const EnumCodeInspectStatus4 EnumCode = 4
		
// InspectStatus + 质检中
const EnumCodeInspectStatus5 EnumCode = 5
		
// InspectStatus + 待出库
const EnumCodeInspectStatus6 EnumCode = 6
		
// InspectStatus + 出库异常
const EnumCodeInspectStatus7 EnumCode = 7
		
// InspectStatus + 待自提
const EnumCodeInspectStatus8 EnumCode = 8
		
// InspectStatus + 已取消已自提
const EnumCodeInspectStatus10 EnumCode = 10
		
// InspectStatus + 已发货
const EnumCodeInspectStatus11 EnumCode = 11
		
// InspectStatus + 待重新送检
const EnumCodeInspectStatus12 EnumCode = 12
		
// InspectStatus + 已达送检上限
const EnumCodeInspectStatus13 EnumCode = 13
		
// InspectStatus + 待驿站入库
const EnumCodeInspectStatus14 EnumCode = 14
		
// OrderScene + 其他
const EnumCodeOrderScene1 EnumCode = 1
		
// OrderScene + 直播间
const EnumCodeOrderScene2 EnumCode = 2
		
// OrderScene + 短视频
const EnumCodeOrderScene3 EnumCode = 3
		
// OrderScene + 商品分享
const EnumCodeOrderScene4 EnumCode = 4
		
// OrderScene + 商品橱窗主页
const EnumCodeOrderScene5 EnumCode = 5
		
// OrderScene + 公众号文章商品卡片
const EnumCodeOrderScene6 EnumCode = 6
		