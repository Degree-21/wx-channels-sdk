package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/iancoleman/strcase"
	"github.com/zsmhub/wx-channels-sdk/apis"
	"github.com/zsmhub/wx-channels-sdk/generate/tool"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
)

// 生成API代码

type Field struct {
	Name string
}

type Api struct {
	IsGet      bool
	DocURL     string
	Name       string
	StructName string
	Method     string
	ApiURL     string
	ReqJson    string
	ReqCode    string
	RespJson   string
	RespCode   string
}

type ApiErrCode struct {
	Code apis.ErrCode `json:"code"`
	Msg  string       `json:"msg"`
}

// 枚举值
type EnumCode struct {
	Code apis.ErrCode `json:"code"`
	Msg  string       `json:"msg"`
	Name string       `json:"name"`
}

var apiDocVar = flag.String("doc", "", "[必填]微信文档地址")
var apiPrefixVar = flag.String("prefix", "", "[选填]生成的文件名前缀")

func main() {
	flag.Parse()

	var docURL, savePath, filePrefix string

	if apiDocVar != nil {
		docURL = *apiDocVar
	}
	if docURL == "" {
		fmt.Println("请输入参数doc(微信文档地址):")
		_, _ = fmt.Scanf("%s", &docURL)
	}
	if docURL == "" {
		tool.Die("必传参数 doc=?")
	}

	if apiPrefixVar != nil {
		filePrefix = *apiPrefixVar
	}

	doc := tool.GetDoc(docURL)

	titleHtml, err := doc.Find("title").Html()
	if err != nil {
		tool.Die("failed to get html: %+v\n", err)
	}
	titleHtml = titleHtml[:strings.Index(titleHtml, " ")]
	savePath = fmt.Sprintf("./apis/%s.go", titleHtml)
	if filePrefix != "" {
		savePath = fmt.Sprintf("./apis/%s-%s.go", filePrefix, titleHtml)
	}
	fmt.Printf("开始抓取和生成API代码，文档地址: %s\n", docURL)

	var api Api
	api.DocURL = docURL
	api.Name = titleHtml
	apiURL := strings.ReplaceAll(doc.Find("pre[class=language-text] > code").Eq(0).Text(), "\n", "")
	if apiURL == "" {
		tool.Die("获取不到接口请求地址")
	}
	apiURL = strings.ReplaceAll(apiURL, "http请求方式：", "")
	apiURLSlice := strings.Split(apiURL, " ")
	if len(apiURLSlice) < 2 {
		for _, v := range []string{"POST", "GET", "PUT", "DELETE"} {
			apiURL = strings.ReplaceAll(apiURL, v, v+" ")
			break
		}
		apiURLSlice = strings.Split(apiURL, " ")
	}
	if len(apiURLSlice) < 2 {
		tool.Die("接口调用请求说明：文档不规范，代码需要兼容下~")
	}
	api.Method = strcase.ToCamel(strings.ToLower(apiURLSlice[0]))
	api.ApiURL = apiURLSlice[1]
	api.ApiURL = strings.ReplaceAll(api.ApiURL[:strings.Index(api.ApiURL, "?")], apis.DefaultWXAPIHost, "")
	if strings.ToUpper(api.Method) != http.MethodPost {
		api.IsGet = true
		api.Method = strcase.ToCamel(strings.ToLower(http.MethodGet))
	}

	// tip: Get 方式的接口量少没做兼容，请求参数需手动整理到 Req 结构体，Post 则不用

	if strings.Contains(api.DocURL, "get.html") {
		fmt.Println("该文档为get文档 Get 方式的接口量少没做兼容，请求参数需手动整理到 Req 结构体，Post 则不用")
	}
	jsonCodeTotal := doc.Find("pre[class=language-json] > code").Length()
	doc.Find("pre[class=language-json] > code").Each(func(i int, selection *goquery.Selection) {
		if i == 0 && (api.IsGet || jsonCodeTotal < 2) {
			api.RespJson = selection.Text()
		} else if i == 0 {
			api.ReqJson = selection.Text()
		} else if i == 1 {
			api.RespJson = selection.Text()
		}
	})

	result, err := generateApiCode(api)
	if err != nil {
		tool.Die("generateApiCode failed: %+v\n", err)
	}
	err = ioutil.WriteFile(savePath, result, os.ModePerm)
	if err != nil {
		tool.Die("ioutil.WriteFile failed: %+v\n", err)
	}
	fmt.Printf("保存文件成功，文件路径: %s\n", savePath)

	addApiErrorCodeToFile(generateApiErrCode(doc))
	//addEnumCodeToFile(generateApiEnumCode(doc))
}

func generateApiCode(api Api) (result []byte, err error) {
	tpl, err := template.ParseFiles("./generate/template/api.tmpl")
	if err != nil {
		return
	}

	segs := strings.Split(api.ApiURL, "/")
	for k, v := range segs {
		if k <= 2 {
			continue
		}
		api.StructName += strcase.ToCamel(v)
	}

	api.ReqCode, err = tool.GenerateStruct(api.ReqJson, "Req"+api.StructName, false)
	if err != nil {
		fmt.Printf("generate reqStruct failed: %+v\n", err)
		return
	}

	api.RespCode, err = tool.GenerateStruct(api.RespJson, "Resp"+api.StructName, false)
	if err != nil {
		fmt.Printf("generate RespStruct failed: %+v\n", err)
		return
	}

	buf := bytes.NewBufferString("")
	err = tpl.Execute(buf, api)
	if err != nil {
		return
	}

	result = buf.Bytes()
	return
}

func generateApiErrCode(doc *goquery.Document) []ApiErrCode {
	var codeSlice []ApiErrCode
	var docErrCode *goquery.Selection
	if doc.Find("#错误码~.table-wrp").Length() > 0 {
		docErrCode = doc.Find("#错误码~.table-wrp")
	} else if doc.Find("#返回码~.table-wrp").Length() > 0 {
		docErrCode = doc.Find("#返回码~.table-wrp")
	} else {
		return codeSlice
	}

	docErrCode.Each(func(i int, selection *goquery.Selection) {
		htmlStr, _ := selection.Find("tr").Html()
		if strings.Contains(htmlStr, "<th style=\"text-align:left\">枚举值</th> <th style=\"text-align:left\">描述</th>") {
			return
		}
		selection.Find("tr").Each(func(i int, selection *goquery.Selection) {
			if i == 0 { // 表头
				return
			}
			var tdSlice []string
			selection.Find("td").Each(func(i int, selection *goquery.Selection) {
				tdSlice = append(tdSlice, selection.Text())
			})
			if len(tdSlice) < 2 { // 参数一般为 2 行
				return
			}
			//fmt.Println(tdSlice)
			codeInt, _ := strconv.Atoi(tdSlice[0])
			code := apis.ErrCode(codeInt)
			if code == apis.ErrCodeSuccess || code == apis.ErrCodeSysErr || code == apis.ErrCodeMinus2 {
				return
			}
			codeSlice = append(codeSlice, ApiErrCode{
				Code: code,
				Msg:  tdSlice[1],
			})
		})
	})
	return codeSlice
}

func generateApiEnumCode(doc *goquery.Document) []EnumCode {
	//enumData := make(map[string]map[string]string)
	if doc.Find("#枚举值~.table-wrp").Length() <= 0 {
		return []EnumCode{}
	}
	//var docErrCode *goquery.Selection
	var codeSlice []EnumCode
	seen := make(map[string]bool)
	doc.Find("h4").Each(func(_ int, h4 *goquery.Selection) {
		id, exists := h4.Attr("id")
		if !exists {
			return
		}

		table := h4.NextAllFiltered("div.table-wrp").First().Find("table")
		if table.Length() == 0 {
			return
		}
		// 检查表头是否包含指定内容
		table.Find("tbody tr").Each(func(_ int, tr *goquery.Selection) {
			if strings.Contains(table.Text(), "字段名") {
				return
			}
			tds := tr.Find("td")
			if tds.Length() < 2 {
				return
			}

			key := tds.Eq(0).Text()
			value := tds.Eq(1).Text()

			codeInt, _ := strconv.Atoi(key)
			code := apis.ErrCode(codeInt)
			codeValue := fmt.Sprintf("%s:%s:%s", key, value, id)
			if !seen[codeValue] {
				codeSlice = append(codeSlice, EnumCode{
					Code: code,
					Msg:  value,
					Name: id,
				})
				seen[codeValue] = true
			}
			return
		})
	})
	return codeSlice
}

func addApiErrorCodeToFile(codes []ApiErrCode) {
	if len(codes) == 0 {
		fmt.Printf("共新增 0 个错误码\n")
		return
	}
	filename := "./apis/api_error.go"
	fileContent := tool.ReadFile(filename)
	var writeTotal int
	for _, v := range codes {
		// 已存在则跳过
		if strings.Contains(fileContent, fmt.Sprintf("const ErrCode%d ErrCode", v.Code)) {
			continue
		}

		content := `
// %s
const ErrCode%d ErrCode = %d
`
		tool.AddContentToFile(filename, fmt.Sprintf(content, v.Msg, v.Code, v.Code))
		writeTotal++
	}

	fmt.Printf("共新增 %d 个错误码\n", writeTotal)
}
func addEnumCodeToFile(codes []EnumCode) {
	if len(codes) == 0 {
		fmt.Printf("共新增 0 个枚举值\n")
		return
	}
	filename := "./apis/api_enum.go"
	fileContent := tool.ReadFile(filename)
	var writeTotal int
	for _, v := range codes {
		//const EnumCodeOrderScene1 EnumCode = 1
		fmt.Println(fmt.Sprintf("const EnumCode%s%v EnumCode", v.Name, v.Code))

		// 已存在则跳过
		if strings.Contains(fileContent, fmt.Sprintf("const EnumCode%s%v EnumCode", v.Name, v.Code)) {
			continue
		}
		//
		content := `
// %s + %s
const EnumCode%s%v EnumCode = %d
		`
		tool.AddContentToFile(filename, fmt.Sprintf(content, v.Name, v.Msg, v.Name, v.Code, v.Code))
		writeTotal++
	}

	fmt.Printf("共新增 %d 个枚举值\n", writeTotal)
}
