package demo

import (
	"fmt"
	"github.com/labstack/echo/v4"
	channels "github.com/zsmhub/wx-channels-sdk"
	"github.com/zsmhub/wx-channels-sdk/callbacks"
)

// 接收视频号回调事件示例
func CallbackMain() {

	client, err := CallbackRepo.InitCallbackHandler()

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	e := echo.New()

	// 接收视频号小店回调事件
	e.GET("/callback/shop", func(c echo.Context) error {
		client.ShopCallback.EchoTestHandler(c.Response().Writer, c.Request())
		return nil
	})
	e.POST("/callback/shop", func(c echo.Context) error {
		// 打印微信回调事件请求体，便于调试
		// requestBody, _ := ioutil.ReadAll(c.Request().Body)
		// fmt.Printf("channels post : uri=%s, req=%s\n", c.Request().RequestURI, string(requestBody))

		_ = CallbackRepo.HandleShopPostReqShouest(c)
		return c.String(200, "success")
	})

	// 接收视频号橱窗回调事件
	e.GET("/callback/window", func(c echo.Context) error {
		client.WindowCallback.EchoTestHandler(c.Response().Writer, c.Request())
		return nil
	})
	e.POST("/callback/window", func(c echo.Context) error {
		// 打印微信回调事件请求体，便于调试
		// requestBody, _ := ioutil.ReadAll(c.Request().Body)
		// fmt.Printf("channels post : uri=%s, req=%s\n", c.Request().RequestURI, string(requestBody))

		_ = CallbackRepo.HandleWindowPostRequest(c)
		return c.String(200, "success")
	})

	e.Logger.Fatal(e.Start(":1323"))
	select {}
}

type callbackRepo struct{}

var CallbackRepo = new(callbackRepo)

// 回调事件初始化
func (callbackRepo) InitCallbackHandler() (*channels.Sdk, error) {
	// 视频号小店回调事件解析
	client := channels.NewSdk()
	if err := client.NewShopCallbackHandler(ShopCallbackToken, ShopCallbackEncodingAESKey); err != nil {
		return nil, err
	}

	// 视频号橱窗回调事件解析
	if err := client.NewWindowCallbackHandler(WindowCallbackToken, WindowCallbackEncodingAESKey); err != nil {
		return nil, err
	}

	return nil, nil
}

// 处理视频号小店回调事件
func (callbackRepo) HandleShopPostReqShouest(c echo.Context) error {
	msg, err := channels.Sdk.ShopCallback.GetCallbackMsg(c.Request())
	if err != nil {
		return err
	}

	go func() {
		defer func() {
			_ = recover()
		}()

		switch msg.MsgType {

		case callbacks.MessageTypeEvent:
			switch msg.EventType {

			case callbacks.EventTypeProductSpuAudit:
				extras, ok := msg.Extras.(callbacks.EventProductSpuAudit)
				if !ok {
					fmt.Println("获取商品审核回调参数失败")
					return
				}
				fmt.Printf("接收到商品审核事件的原始内容：%s\n", msg.OriginalMessage)
				fmt.Printf("解析商品审核事件的结构体：%+v\n", extras)
			}

		}
	}()

	return nil
}

// 处理视频号橱窗回调事件
func (r callbackRepo) HandleWindowPostRequest(c echo.Context) error {
	msg, err := channels.Sdk.WindowCallback.GetCallbackMsg(c.Request())
	if err != nil {
		return err
	}

	go func() {
		defer func() {
			_ = recover()
		}()

		fmt.Printf("接收到回调事件的原始内容：%s\n", msg.OriginalMessage)
	}()

	return nil
}
