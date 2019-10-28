package exchange_api_demo_golang

import (
	"sort"
	"crypto/md5"
	"strings"
	"encoding/hex"
	"net/http"
	"fmt"
	"io/ioutil"
	"net/url"
)

// 获取签名参数
func GetSignData(data map[string]string, secretKey string) string {

	var signStr string
	tempSlice := make([]string, 0)
	for key := range data {
		tempSlice = append(tempSlice, key)
	}
	sort.Strings(tempSlice)
	for _, v := range tempSlice {
		signStr += v + "=" + data[v] + "&"
	}
	signStr = signStr + "secret_key=" + secretKey

	hash := md5.Sum([]byte(signStr))
	hashed := hash[:]

	return strings.ToUpper(hex.EncodeToString(hashed))
}

// Http Get请求基础函数, 通过封装Go语言Http请求, 支持ZT网REST API的HTTP Get请求
// strUrl: 请求的URL
// return: 请求结果
func HttpGetRequest(strUrl string) (string, error) {

	req, err := http.NewRequest("GET", strUrl, nil)
	if err != nil {
		fmt.Println(err)
		return err.Error(), err
	}
	req.Header.Add("X-SITE-ID", X_SITE_ID)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error(), err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error(), err
	}

	return string(body), nil
}

// Http POST请求基础函数, 通过封装Go语言Http请求, 支持ZT网REST API的HTTP POST请求
// strUrl: 请求的URL
// return: 请求结果
func HttpPostRequest(mapParams map[string]string, strUrl string) (string, error) {

	mapParams["api_key"] = API_KEY
	sign := GetSignData(mapParams, SECRET_KEY)
	mapParams["sign"] = sign

	u := url.Values{}
	for k, v := range mapParams {
		u.Set(k, v)
	}
	rd := strings.NewReader(u.Encode())

	req, err := http.NewRequest("POST", strUrl, rd)
	req.Header.Add("X-SITE-ID", X_SITE_ID)
	req.Header.Add("Content-type", ContentType)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error(), err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error(), err
	}

	return string(body), nil
}

