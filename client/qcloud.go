package client

import (
	"sort"
	"strings"
	"fmt"
	"encoding/base64"
	"crypto/hmac"
	"crypto/sha256"
	"github.com/parnurzeal/gorequest"
	"encoding/json"
	"github.com/pkg/errors"
	"qiniupkg.com/x/url.v7"
)

type PublicVar struct {
	Action          string
	Region          string
	Timestamp       string
	Nonce           string
	SignatureMethod string
	SecretId        string
}
type SslVar struct {
	Page string
	//Count     int
	//searchKey string
	//certType  string
	//id        string
	//withCert  int
}
type SslInfo struct {
	Code     int      `json:"code"`
	Message  string   `json:"message"`
	CodeDesc string   `json:"codeDesc"`
	Data     DataInfo `json:"data"`
}
type DataInfo struct {
	TotalNum int           `json:"totalNum"`
	List     []SslInfoList `json:"list"`
}
type SslInfoList struct {
	OwnerUin            string `json:"ownerUin"`
	ProjectId           string `json:"projectId"`
	From                string `json:"from"`
	Type                string `json:"type"`
	CertType            string `json:"certType"`
	ProductZhName       string `json:"productZhName"`
	Domain              string `json:"domain"`
	Alias               string `json:"alias"`
	Status              int    `json:"status"`
	extra               string `json:"extra"`
	VulnerabilityStatus string `json:"vulnerability_status"`
	StatusMsg           string `json:"statusMsg"`
	VerifyType          string `json:"verifyType"`
	CertBeginTime       string `json:"certBeginTime"`
	CertEndTime         string `json:"certEndTime"`
	ValidityPeriod      string `json:"validityPeriod"`
	InsertTime          string `json:"insertTime"`
	ProjectInfo struct {
		ProjectId  string `json:"projectId"`
		OwnerUin   int    `json:"ownerUin"`
		Name       string `json:"name"`
		CreatorUin int    `json:"creatorUin"`
		CreateTime string `json:"createTime"`
		Info       string `json:"info"`
	} `json:"projectInfo"`
	Id              string   `json:"id"`
	SubjectAltName  []string `json:"subjectAltName"`
	TypeName        string   `json:"type_name"`
	StatusName      string   `json:"status_name"`
	IsVip           bool     `json:"is_vip"`
	IsDv            bool     `json:"is_dv"`
	IsWildcard      bool     `json:"is_wildcard"`
	IsVulnerability bool     `json:"is_vulnerability"`
	RenewAble       bool     `json:"renew_able"`
}

//type ProjectInfoInfo struct {
//	ProjectId  string `json:"projectId"`
//	OwnerUin   int64  `json:"ownerUin"`
//	Name       string `json:"name"`
//	CreatorUin int64  `json:"creatorUin"`
//	CreateTime string `json:"createTime"`
//	Info       string `json:"info"`
//}

const (
	SecretId  string = "AKID34SsTvKQQFuOLHpnKgs5GNakv2FTBaUZ"
	API       string = "wss.api.qcloud.com/v2/index.php?"
	SecretKey string = "Rn17NZdZYmwyqqoeYPAC8vXyFLG9ypv9"
)

func GetSslInfo(publicVar PublicVar, sslVar SslVar) (string, error) {
	varMap := make(map[string]string)
	varMap["Action"] = publicVar.Action
	varMap["Region"] = publicVar.Region
	varMap["Timestamp"] = string(publicVar.Timestamp)
	varMap["Nonce"] = string(publicVar.Nonce)
	varMap["SignatureMethod"] = publicVar.SignatureMethod
	varMap["SecretId"] = SecretId
	varMap["Page"] = string(sslVar.Page)
	//varMap["Count"] = string(sslVar.Count)
	//varMap["searchKey"] = sslVar.searchKey
	//varMap["certType"] = sslVar.certType
	//varMap["id"] = sslVar.id
	//varMap["withCert"] = string(sslVar.withCert)
	var varSlice []string
	for key, _ := range varMap {
		varSlice = append(varSlice, key)
	}
	sort.Strings(varSlice)
	//sort
	var str string
	for i := 0; i < len(varSlice); i++ {
		str += fmt.Sprintf("%s", varSlice[i]+"="+varMap[varSlice[i]]+"&")
	}

	//get Signature
	pinStr := "GET" + API + strings.Trim(str, "&")
	signByte := []byte{}
	temp := hmac.New(sha256.New, []byte(SecretKey))
	temp.Write([]byte(pinStr))
	signByte = temp.Sum(nil)
	signStr := base64.StdEncoding.EncodeToString(signByte)
	signStr = url.QueryEscape(signStr)

	//get request
	requestUrl := "https://" + API + "&" + str + "Signature=" + signStr
	//fmt.Println(requestUrl)
	request := gorequest.New()
	var sslInfo SslInfo
	resp, body, errs := request.Get(requestUrl).End()
	if resp.StatusCode != 200 || len(errs) != 0 {
		newError := errors.New("resp is not 200 or errs is not null")
		return "", newError
	}
	sslInfo.Data.List = make([]SslInfoList, 10)
	//fmt.Println("-------------------")
	json.Unmarshal([]byte(body), &sslInfo)
	//fmt.Println(sslInfo.Code)
	//fmt.Println(sslInfo.CodeDesc)

	//return sslEndTime
	var sslEndTime string
	for _, val := range sslInfo.Data.List {
		sslEndTime += fmt.Sprintf("ssl cert is %s and the EndTime is %s\n", val.Alias, val.CertEndTime)
	}
	if sslInfo.CodeDesc != "Success" || sslInfo.Code != 0 {
		newError := errors.New("resp is 200 but request is wrong")
		return "", newError
	}

	return sslEndTime, nil

}
