package sys

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SignRobot(urlAddress, Secret string) string {
	//  构建 签名
	//  把timestamp+"\n"+密钥当做签名字符串，使用HmacSHA256算法计算签名，然后进行Base64 encode，最后再把签名参数再进行urlEncode，得到最终的签名（需要使用UTF-8字符集）。
	timeStampNow := time.Now().UnixNano() / 1000000
	signStr := fmt.Sprintf("%d\n%s", timeStampNow, Secret)

	hash := hmac.New(sha256.New, []byte(Secret))
	hash.Write([]byte(signStr))
	sum := hash.Sum(nil)

	encode := base64.StdEncoding.EncodeToString(sum)
	urlEncode := url.QueryEscape(encode)

	// 构建 请求 url
	UrlAddress := fmt.Sprintf("%s&timestamp=%d&sign=%s", urlAddress, timeStampNow, urlEncode)
	return UrlAddress
}

// SendMsg 发送消息基础模版
func SendMsg(urlAddress, secret, params string) string {
	client := &http.Client{}
	UrlAddress := SignRobot(urlAddress, secret)
	req, err := http.NewRequest("POST", UrlAddress, strings.NewReader(params))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}
