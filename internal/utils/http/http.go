package http

import (
	useragent "crawler/internal/utils/user_agent"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

var (
	httpClient *http.Client
	clientOnce sync.Once
)

// 创建 http 连接池
func getHTTPClient() *http.Client {
	clientOnce.Do(func() {
		transport := &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			MaxIdleConnsPerHost:   100,
			MaxConnsPerHost:       100,
		}
		// 解析代理服务器地址
		// proxy, err := url.Parse(configs.Cfg.Proxy)
		// if err != nil {
		// 	fmt.Println("Error parsing proxy URL:", err)
		// } else {
		// 	transport.Proxy = http.ProxyURL(proxy)
		// }
		httpClient = &http.Client{
			Transport: transport,
			Timeout:   30 * time.Second,
		}
	})
	return httpClient
}

// HttpGet 发送 GET 请求
func HttpGet(url string) (body []byte) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("%s构建请求失败: %s \n", url, err)
		return
	}
	// 设置请求头，指定接受 UTF-8 编码的响应
	req.Header.Set("Accept-Charset", "utf-8")
	req.Header.Set("User-Agent", useragent.RandomUserAgent())
	// 发送请求
	client := getHTTPClient()
	for i := 0; i < 3; i++ {
		resp, err := client.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			// 读取响应体
			body, err = io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("%s 读取响应体失败: %s \n", url, err)
			} else {
				// 响应内容
				return
			}
		}
		time.Sleep(time.Second * 1)
	}
	return
}

// HttpGet Json 转换
func HttpToJson(url string) (data string) {
	body := HttpGet(url)
	if len(body) == 0 {
		return
	}
	return EscapeUnicodeInJSON(body)
}

// EscapeUnicodeInJSON 将 JSON 字符串中的 Unicode 编码字符串转义成正常的中文字符串
func EscapeUnicodeInJSON(jsonStr []byte) string {
	// 使用 json.Unmarshal 解析 JSON 字符串
	var data interface{}
	if err := json.Unmarshal(jsonStr, &data); err != nil {
		fmt.Println("解析 JSON 字符串失败:", err)
		return ""
	}

	// 使用 json.Marshal 转换为带转义的 JSON 字符串
	escapedJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println("转换为 JSON 字符串失败:", err)
		return ""
	}

	return string(escapedJSON)
}
