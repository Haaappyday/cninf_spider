package parse

import (
	"encoding/json"
	"io"
	"net/http"
)

// 巨潮公告resp结构
type Announcement struct {
	Id                    string
	SecCode               string
	SecName               string
	OrgId                 string
	AnnouncementId        string
	AnnouncementTitle     string
	AnnouncementTime      int64
	AdjunctUrl            string
	AdjunctSize           int64
	AdjunctType           string
	StorageTime           string
	ColumnId              string
	PageColumn            string
	AnnouncementType      string
	AssociateAnnouncement string
	Important             string
	BatchNum              string
	AnnouncementContent   string
	OrgName               string
	TileSecName           string
	ShortTitle            string
	AnnouncementTypeName  string
	SecNameList           string
}

type RespBody struct {
	ClassifiedAnnouncements string
	TotalSecurities         int
	TotalAnnouncement       int
	TotalRecordNum          int
	Announcements           []Announcement
	CategoryList            string
	HasMore                 bool
	Totalpages              int
}

// 一下是一个返回体的示例,最终要解析到上述的struct中
// {
// 	"classifiedAnnouncements": null,
//     "totalSecurities": 0,
//     "totalAnnouncement": 528,
//     "totalRecordNum": 528,
//     "announcements": [
//         {
//             "id": null,
//             "secCode": "002104",
//             "secName": "恒宝股份",
//             "orgId": "9900001841",
//             "announcementId": "1217805900",
//             "announcementTitle": "关于收到《行政<em>处罚</em>事先告知书》的公告",
//             "announcementTime": 1694102400000,
//             "adjunctUrl": "finalpage/2023-09-08/1217805900.PDF",
//             "adjunctSize": 275,
//             "adjunctType": "PDF",
//             "storageTime": null,
//             "columnId": "09020202||250101||251302",
//             "pageColumn": "SZZB",
//             "announcementType": "01010501||010112||012399",
//             "associateAnnouncement": null,
//             "important": null,
//             "batchNum": null,
//             "announcementContent": "",
//             "orgName": null,
//             "tileSecName": "恒宝股份",
//             "shortTitle": "关于收到《行政<em>处罚</em>事先告知书》的公告",
//             "announcementTypeName": null,
//             "secNameList": null
//         },
// 	]
// 	"categoryList": null,
//     "hasMore": true,
//     "totalpages": 17
// }

func ParseResp(resp *http.Response) (*RespBody, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rb RespBody
	err = json.Unmarshal([]byte(string(body)), &rb)
	if err != nil {
		return nil, err
	}
	return &rb, nil

}
