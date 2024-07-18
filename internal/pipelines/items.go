package pipelines

import (
	"crawler/global"
	"crawler/internal/utils"
	"errors"
	"fmt"
	"strings"
)

type ItemVideo struct {
	Playlink string
	Name     string
}

// 处理视频 Url 分割
func processItemPlayUrl(playUrl, flags string) (item ItemVideo, err error) {
	playUrl = strings.TrimSpace(playUrl)
	var playData []string
	switch flags {
	case global.Yzzy_Source: // 处理 playUrl
		if strings.Contains(playUrl, "$$") {
			playUrl = strings.ReplaceAll(playUrl, "$$", "$")
		}
		if strings.Contains(playUrl, "https") && strings.Contains(playUrl, "$") {
			playData = strings.Split(playUrl, "$")
			if len(playData) == 1 {
				fmt.Println("vodPlayUrl:", playUrl)
				return item, errors.New("vodPlayUrl")
			}
			item.Playlink = playData[1]
			item.Name = playData[0]
		} else {
			playData = strings.Split(playUrl, ":")
			if len(playData) > 1 {
				item.Playlink = "https:" + playData[1]
				item.Name = playData[0]
			}
		}
	case global.Ffzy_Source:
		if !strings.Contains(playUrl, "https") {
			fmt.Println("vodPlayUrl 未包含https:", playUrl)
			return item, errors.New("vodPlayUrl 未包含$")
		}
		playData = strings.Split(playUrl, "$")
		if !strings.Contains(playUrl, "$") {
			playData = strings.Split(playUrl, "http")
		}
		item.Playlink = playData[1]
		item.Name = playData[0]
	default:
		err = fmt.Errorf("flags error: %s", flags)
	}
	return
}

// 处理视频 Url 分割
func processItemVodPlayUrl(VodPlayUrl, delimiter, flags string) (playList []string) {
	if VodPlayUrl == "" {
		return
	}
	switch flags {
	case global.Yzzy_Source:
		if !strings.Contains(VodPlayUrl, "#") {
			playList = []string{VodPlayUrl}
		} else {
			playList = strings.Split(VodPlayUrl, "#")
		}
	case global.Ffzy_Source:
		if len(delimiter) > 3 {
			delimiter = delimiter[:3]
		}
		if !strings.Contains(VodPlayUrl, delimiter) {
			fmt.Println(VodPlayUrl, "not contains ", delimiter)
		} else {
			vodPlayUrl := strings.Split(VodPlayUrl, delimiter)
			if len(vodPlayUrl) > 1 {
				playUrl := vodPlayUrl[1]
				if strings.Contains(playUrl, "https") && strings.Contains(playUrl, "#") {
					vodPlayUrl = strings.Split(playUrl, "#")
				}
				for _, v := range vodPlayUrl {
					if strings.Contains(v, "m3u8") {
						playList = append(playList, v)
					}
				}
			} else {
				fmt.Println("processItemVodPlayUrl error")
			}
		}
	}
	return
}

// 判断专辑是否完结
func processItemVodStatus(VodRemarks string, total, VideoType int) (EpisodeStatus, Episodes int) {
	if strings.Contains(VodRemarks, "全") {
		EpisodeStatus = 1
		Episodes = total
	} else if strings.Contains(VodRemarks, "更新") {
		EpisodeStatus = 0
		Episodes = total
	} else if VideoType == 2 {
		EpisodeStatus = 1
		Episodes = total
	} else {
		EpisodeStatus = -1
		Episodes = 0
	}
	return
}

// 判断类型
func processItemVodType(vodType interface{}, flags string) int {
	//  1:电视剧 2: 电影 3:动漫  4:综艺    100:短剧
	vt := utils.AssertTypeIdType(vodType)
	if flags == global.Ffzy_Source {
		switch vt {
		case "13", "14", "15", "16", "21", "23": // 电视剧
			return 1
		case "6", "7", "8", "9", "10", "11", "12", "20", "34": // 电影
			return 2
		case "29": // 动漫
			return 3
		case "25": // 综艺
			return 4
		case "36": // 短剧
			return 100
		default: // 其他
			return -1
		}
	} else {
		switch vt {
		case "12", "14", "15": // 电视剧
			return 1
		case "5", "6", "7", "8", "9", "10", "11", "41": // 电影
			return 2
		case "66", "67": //动漫
			return 3
		case "62": //综艺
			return 4
		case "83": //短剧
			return 100
		default: // 其他
			return -1
		}
	}
}

// 处理 source 字段
func processItemSource(source string) string {
	if strings.Contains(source, global.Ffzy) {
		return global.Ffzy_Source
	}
	if strings.Contains(source, global.Yzzy) {
		return global.Yzzy_Source
	}
	return "default"
}
