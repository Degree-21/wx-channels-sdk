package tool

import (
	"fmt"
	"github.com/ChimeraCoder/gojson"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func GetDoc(docURL string) *goquery.Document {
	resp, err := http.Get(docURL)
	if err != nil {
		Die("http get of errcode documentation failed: %+v\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		Die("non-200 app: %+v\n", resp)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		Die("parse document failed: %+v\n", err)
	}
	return doc
}

func Die(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

func GenerateStruct(rawJson string, structName string, subStruct bool) (code string, err error) {
	if rawJson == "" {
		code = fmt.Sprintf("type %s struct{}", structName)
		return
	}
	var rawCode []byte
	rawCode, err = gojson.Generate(strings.NewReader(rawJson),
		gojson.ParseJson,
		structName,
		"apis",
		[]string{"json"}, subStruct, true,
	)
	if err != nil {

		// 找到最后一个逗号的位置
		lastCommaIndex := strings.LastIndex(rawJson, ",")
		if lastCommaIndex == -1 {
			// 如果没有逗号，则返回原始的 JSON 字符串
			//return jsonStr
			fmt.Println(err.Error())
			return
		}
		remove := rawJson[:lastCommaIndex] + rawJson[lastCommaIndex+1:]
		rawCode, err = gojson.Generate(strings.NewReader(remove),
			gojson.ParseJson,
			structName,
			"apis",
			[]string{"json"}, subStruct, true,
		)
		if err != nil {
			fmt.Println("=====")
			fmt.Println(err.Error())
			fmt.Println(remove)
			return
		}
		//
		//var t interface {
		//}
		//err = json.Unmarshal([]byte(remove), &t)
		//if err != nil {
		//	fmt.Println(err.Error())
		//	return
		//}
		//
		//// 移除最后一个逗号及其后面的所有字符
		////return
		//
		////fmt.Println("1111")
		////fmt.Println(rawJson)
		////fmt.Println(err.Error())
		//return
	}

	code = string(rawCode)
	code = strings.Replace(code, "package apis\n\n", "", 1)
	code = strings.ReplaceAll(code, "int64", "int")
	commonRespRegexp := regexp.MustCompile("Errcode.+\\n.+?Errmsg.+")
	code = commonRespRegexp.ReplaceAllString(code, "CommonResp")
	return
}

// 添加内容到文件中
func AddContentToFile(filename, content string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		Die("err=%+v", err)
	}
	defer func() {
		_ = file.Close()
	}()
	if _, err := file.WriteString(content); err != nil {
		Die("err=%+v", err)
	}
}

// 读取整个文件
func ReadFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		Die("err=%+v", err)
	}
	return string(content)
}
