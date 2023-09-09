package main

import (
	"bytes"
	"cinef_spider/define"
	"cinef_spider/parse"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return text
}

func inputNone(key string) string {
	if key == "none" {
		return ""
	} else {
		return key
	}
}

func main() {
	var stock, searchkey, category, trade, seDate string

	stock = input("输入股票代码(示例:000001,gssz0000001),空填none:")
	searchkey = input("输入标题关键词,多条件以;分隔,空填none:")
	category = input("输入报告类型,详情见README,空填none:")
	trade = input("输入行业类型,详情见README,空填none:")
	seDate = input("输入日期范围(示例:2021-01-01~2023-08-31),不可为空:")

	stock = inputNone(stock)
	searchkey = inputNone(searchkey)
	category = inputNone(category)
	trade = inputNone(trade)
	if seDate == "" {
		panic("日期不可为空")
	}

	// stock, searchkey, category, trade, seDate = "", "处罚", "", "采矿业", "2021-01-01~2023-08-31"
	fileDown(stock, searchkey, category, trade, seDate)
}

// 1, "", "处罚", "", "采矿业", "2021-01-01~2023-08-31"

func getParse(reqReader *bytes.Reader) (*parse.RespBody, error) {
	req, err := http.NewRequest("POST", define.URL, reqReader)
	if err != nil {
		return nil, err
	}

	req.Header = map[string][]string{
		"User-Agent":   {define.UserAgent},
		"Content-Type": {define.ContentType},
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("http status error")
	}
	defer resp.Body.Close()

	rb, err := parse.ParseResp(resp)
	if err != nil {
		return nil, err
	}
	return rb, nil
}

type baseInfo struct {
	totalReportNum int
	totalPageNum   int
	lastReportNum  int
}

func getBaseInfo(stock, searchkey, category, trade, seDate string) (*baseInfo, error) {
	reqReader := define.ReqReader(1, stock, searchkey, category, trade, seDate)
	rb, err := getParse(reqReader)
	if err != nil {
		return nil, err
	}

	var lastreportNum int
	if rb.TotalRecordNum-rb.Totalpages*30 <= 0 {
		lastreportNum = rb.TotalRecordNum
	} else {
		lastreportNum = rb.TotalRecordNum - rb.Totalpages*30
	}

	bi := &baseInfo{
		totalReportNum: rb.TotalRecordNum, // 全部报告数量
		totalPageNum:   rb.Totalpages + 1, // 合计页数, 返回值需要+1
		lastReportNum:  lastreportNum,     // 最后一页报告数量
	}
	return bi, nil
}

func fileDown(stock, searchkey, category, trade, seDate string) {
	// 1, "", "处罚", "", "采矿业", "2021-01-01~2023-08-31"
	info, err := getBaseInfo("", "处罚", "", "采矿业", "2021-01-01~2023-08-31")
	if err != nil {
		fmt.Printf("get info err: %v\n", err)
	}

	totalPageNum := info.totalPageNum                        // 合计页数, 返回值需要+1
	lastReportNum := info.lastReportNum                      // 最后一页报告数量
	replacer := strings.NewReplacer("<em>", "", "</em>", "") // 替换<em>标签

	for page := 1; page < totalPageNum+1; page++ {
		reqReader := define.ReqReader(page, stock, searchkey, category, trade, seDate)
		rb, err := getParse(reqReader)
		if err != nil {
			fmt.Printf("parse resp failed, err: %v\n", err)
		}

		if rb.HasMore {
			for i := 0; i < 30; i++ {
				part := rb.Announcements[i]
				title := replacer.Replace(part.AnnouncementTitle)
				date := time.Unix(part.AnnouncementTime/1000, 0).Format("2006-01-02") // 1685635200000 最后三位数nsec
				url := "http://static.cninfo.com.cn/" + part.AdjunctUrl

				err := down(url, date+" "+title+".pdf")
				if err != nil {
					fmt.Printf("file download failed, err: %v\n", err)
				}
			}
		}

		if !rb.HasMore {
			for i := 0; i < lastReportNum; i++ {
				part := rb.Announcements[i]
				title := replacer.Replace(part.AnnouncementTitle)
				date := time.Unix(part.AnnouncementTime/1000, 0).Format("2006-01-02") // 1685635200000 最后三位数nsec
				url := "http://static.cninfo.com.cn/" + part.AdjunctUrl

				err := down(url, date+" "+title+".pdf")
				if err != nil {
					fmt.Printf("file download failed, err: %v\n", err)
				}
			}
		}
	}

}

func down(url string, name string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp == nil {
		return errors.New("response is empty")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("http status error")
	}

	file, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	os.WriteFile("./result/"+name, file, 0644)

	return nil
}
