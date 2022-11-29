package sdk

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	HttpClient *http.Client
	AppID      string
	Secret     string
	BaseUrl    string
}

func NewClient(baseUrl, appID, secret string) (*Client, error) {
	client := new(Client)
	client.HttpClient = &http.Client{
		//Timeout: 500 * time.Millisecond,
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			//AllowHTTP: true,
			//Dial:                  dialer.Dial,
			DisableKeepAlives:     true,
			MaxIdleConnsPerHost:   -1,
			ResponseHeaderTimeout: 10 * time.Second,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		},
	}
	client.AppID = appID
	client.Secret = secret
	client.BaseUrl = baseUrl
	return client, nil
}
func (c *Client) DoPost(addr string, Header map[string]string, body string, jsonFormat bool) (string, error) {
	req, err := http.NewRequest("POST", addr, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	if Header != nil {
		for k, v := range Header {
			req.Header.Set(k, v)
		}
	}
	if jsonFormat {
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	req.Header.Set("Content-Length", strconv.Itoa(len(body)))
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return "", err
	}
	if resp.Body != nil {
		respon, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		resp.Body.Close()
		return string(respon), nil
	}
	return "", nil
}

func (c *Client) SignHelper(appID, secret string, signType uint8, params map[string]interface{}) (string, error) {

	// 对params 进行升序列排序
	keys := []string{}
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStrs := []string{}
	for _, v := range keys {
		signStrs = append(signStrs, fmt.Sprintf("%s=%v", v, params[v]))
	}
	signStr := strings.Join(signStrs, "&")
	dstStr := fmt.Sprintf("%s&%s&%s", appID, secret, signStr)
	if signType == 0 {
		md5Str := fmt.Sprintf("%x", md5.Sum([]byte(dstStr)))
		return md5Str, nil
	} else if signType == 1 {
		shaStr := fmt.Sprintf("%x", sha1.Sum([]byte(dstStr)))
		return shaStr, nil
	}
	return "", errors.New("签名方式不支持")
}

type JSONResponse struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	RS_OK           = "0"
	RS_INTERVAL_ERR = "1"
	RS_DENIED       = "2"
	RS_PARAMS       = "3"
	RS_RPC_ERR      = "4"
)
