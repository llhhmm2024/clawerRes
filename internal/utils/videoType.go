package utils

// 返回 videoType 类型
func VideoType(vodType string) int {
	switch vodType {
	case "5", "6", "7", "8", "9", "10", "11", "41": // 电影
		return 2
	case "12", "14", "15": // 电视剧
		return 1
	case "62": //综艺
		return 3
	case "66", "67": //动漫
		return 4
	case "83": //短剧
		return 100
	default: // 其他
		return -1
	}
}
