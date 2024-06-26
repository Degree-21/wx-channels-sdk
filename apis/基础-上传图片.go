package apis

import (
	"encoding/json"
)

// 上传图片，固定上传类型upload_type=1
// 文档：https://developers.weixin.qq.com/doc/channels/API/basics/img_upload.html

type ReqImgUpload struct {
	// UploadType 上传类型。0: 二进制流；1: 图片 URL（不支持 301/302 跳转）。该参数为 URL 参数
	UploadType int `json:"upload_type" url:"upload_type"`
	// RespType 返回数据类型。0: media_id 和 pay_media_id；1: 图片链接（商品信息相关图片请务必使用此参数得到链接）。该参数为 URL 参数
	RespType int `json:"resp_type" url:"resp_type"`
	// Height 图片的高，单位：像素。upload_type=0 时必填。该参数为 URL 参数
	Height *int `json:"height,omitempty" url:"height,omitempty" `
	// Width 图片的宽，单位：像素。upload_type=0 时必填。该参数为 URL 参数
	Width *int `json:"width,omitempty" url:"width,omitempty"`
	// ImgURL 图片 URL。upload_type=1 时必填。该参数为 POST 请求参数
	ImgURL string `json:"img_url,omitempty" form:"img_url,omitempty" `
}

var _ bodyer = ReqImgUpload{}

func (x ReqImgUpload) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespImgUpload struct {
	CommonResp
	PicFile struct {
		MediaId    string `json:"media_id"`
		PayMediaId string `json:"pay_media_id"`
		ImgUrl     string `json:"img_url"`
	} `json:"pic_file"`
}

var _ bodyer = RespImgUpload{}

func (x RespImgUpload) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecImgUpload(req ReqImgUpload) (RespImgUpload, error) {
	var resp RespImgUpload
	reqUrl := "/channels/ec/basics/img/upload"
	err := c.executeWXApiPost(reqUrl, req, &resp, true)
	if err != nil {
		return RespImgUpload{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespImgUpload{}, bizErr
	}
	return resp, nil
}
