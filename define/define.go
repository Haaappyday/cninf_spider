package define

import (
	"bytes"
	"net/url"
	"strconv"
)

var URL = "http://www.cninfo.com.cn/new/hisAnnouncement/query"
var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"
var ContentType = "application/x-www-form-urlencoded"

// http.NewRequest 的 body 参数
// 一写填写示例：详细信息在README文档中
// pangeNum: 1
// stock: "000001,gssz0000001"
// searchkey: "处罚", 多个关键词中间加;,逻辑为或
// category: "category_bndbg_szsh" ...
// trade: "采矿业" ...
// seDate: "2023-03-09~2023-09-09"
func ReqReader(pageNum int, stock, searchkey, category, trade, seDate string) *bytes.Reader {
	postFormValues := url.Values{
		"pageNum":   {strconv.Itoa(pageNum)},
		"pageSize":  {"30"},
		"column":    {"szse"},
		"tabName":   {"fulltext"},
		"plate":     {},
		"stock":     {stock},
		"searchkey": {searchkey},
		"secid":     {},
		"category":  {category},
		"trade":     {trade},
		"seDate":    {seDate},
		"sortName":  {},
		"sortType":  {},
		"isHLtitle": {"true"},
	}

	postFormDataStr := postFormValues.Encode()
	postFormBytesReader := bytes.NewReader([]byte(postFormDataStr))

	return postFormBytesReader
}
