package utils

import (
	"fmt"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// FormatTag 格式化标签
func FormatTag(v string) string {
	return strings.Join(splitStr(v), ";")
}

// FormatStr 格式化字符串
func FormatStr(v string) string {
	return strings.Join(splitStr(v), ";")
}

func splitStr(s string) []string {
	if strings.Contains(s, ",") {
		return strings.Split(s, ",")
	}
	return []string{s}
}

func GenerateUuid() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

// 获取 Url 目录
func GetParentURL(play_link string) string {
	u := parseUrl(play_link)
	u.Path = path.Dir(u.Path)

	return u.String()
}

func parseUrl(play_link string) *url.URL {
	u, err := url.Parse(play_link)
	if err != nil {
		return nil
	}
	return u
}

// 拼接新的 url path
func JoinNewUrl(play_link string) string {
	u := parseUrl(play_link)
	return u.Path
}

// 获取 scheme 和 path
func GetSchemeAndPath(play_link string) string {
	u := parseUrl(play_link)
	dir := path.Dir(u.Path)
	u.Path = dir
	return u.String()
}

// 拼接 path + filename
func JoinPathFileName(gatherurl, filename string) string {
	if !strings.HasPrefix(gatherurl, "http://") && !strings.HasPrefix(gatherurl, "https://") {
		gatherurl = "https://" + gatherurl // 或者使用 "https://" 根据需要
	}

	// 解析URL
	parsedURL, err := url.Parse(gatherurl)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
	}

	// 拼接路径
	parsedURL.Path = path.Join(parsedURL.Path, filename)

	// 生成完整的URL
	return parsedURL.String()
}

// 是否包含 https
func IsHttps(url string) string {
	if url == "" {
		return ""
	}
	if strings.HasPrefix(url, "https://") {
		return url
	} else {
		var tmp []string
		if strings.Contains(url, ":") {
			tmp = strings.Split(url, ":")
		} else if strings.Contains(url, "//") {
			tmp = strings.Split(url, "//")
		} else {
			return ""
		}
		return "https:" + tmp[1]
	}
}

// 断言类型
func AssertTypeIdType(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case float64, float32:
		return fmt.Sprintf("%.f", v)
	default:
		return ""
	}
}

// 断言 page 类型为 int
func AssertPageType(v interface{}) int {
	switch v := v.(type) {
	case int:
		return v
	case string:
		i, _ := strconv.Atoi(v)
		return i
	default:
		return 1
	}
}
