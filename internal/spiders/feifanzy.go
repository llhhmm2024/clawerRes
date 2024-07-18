package spiders

import (
	"crawler/configs"
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	baseUrl = "https://cj.ffzyapi.com/api.php/provide/vod/?ac=detail&t=%d"
)

// 初始化所有的 Url
func (m *Media) initFeifanzy() {
	ffzyCategory := `[{"type_id":1,"type_pid":0,"type_name":"电影片"},{"type_id":2,"type_pid":0,"type_name":"连续剧"},{"type_id":3,"type_pid":0,"type_name":"综艺片"},{"type_id":4,"type_pid":0,"type_name":"动漫片"},{"type_id":6,"type_pid":1,"type_name":"动作片"},{"type_id":7,"type_pid":1,"type_name":"喜剧片"},{"type_id":8,"type_pid":1,"type_name":"爱情片"},{"type_id":9,"type_pid":1,"type_name":"科幻片"},{"type_id":10,"type_pid":1,"type_name":"恐怖片"},{"type_id":11,"type_pid":1,"type_name":"剧情片"},{"type_id":12,"type_pid":1,"type_name":"战争片"},{"type_id":13,"type_pid":2,"type_name":"国产剧"},{"type_id":14,"type_pid":2,"type_name":"香港剧"},{"type_id":15,"type_pid":2,"type_name":"韩国剧"},{"type_id":16,"type_pid":2,"type_name":"欧美剧"},{"type_id":20,"type_pid":1,"type_name":"记录片"},{"type_id":21,"type_pid":2,"type_name":"台湾剧"},{"type_id":22,"type_pid":2,"type_name":"日本剧"},{"type_id":23,"type_pid":2,"type_name":"海外剧"},{"type_id":24,"type_pid":2,"type_name":"泰国剧"},{"type_id":25,"type_pid":3,"type_name":"大陆综艺"},{"type_id":26,"type_pid":3,"type_name":"港台综艺"},{"type_id":27,"type_pid":3,"type_name":"日韩综艺"},{"type_id":28,"type_pid":3,"type_name":"欧美综艺"},{"type_id":29,"type_pid":4,"type_name":"国产动漫"},{"type_id":30,"type_pid":4,"type_name":"日韩动漫"},{"type_id":31,"type_pid":4,"type_name":"欧美动漫"},{"type_id":32,"type_pid":4,"type_name":"港台动漫"},{"type_id":33,"type_pid":4,"type_name":"海外动漫"},{"type_id":34,"type_pid":1,"type_name":"伦理片"},{"type_id":36,"type_pid":2,"type_name":"短剧"}]`

	// 将 JSON 数据解析到 mediaTypes 切片中
	var mediaTypes []MediaType
	if err := json.Unmarshal([]byte(ffzyCategory), &mediaTypes); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	cfg := configs.Cfg.FfzyList
	filterTypePid := make(map[string]bool, len(cfg))
	for _, v := range cfg {
		filterTypePid[v] = true
	}

	// 要过滤掉的type_id列表
	filterTypeIDs := map[int]bool{
		20: true, 34: true, 22: true, 23: true, 24: true, 16: true, 14: true,
		21: true, 26: true, 27: true, 28: true, 30: true, 31: true, 32: true,
		33: true,
		15: true, 36: true,
	}
	// 根据 type_pid 进行分类
	data := make(map[int][]MediaType)
	for _, mediaType := range mediaTypes {
		// 过滤掉不需要的类型
		if _, ok := filterTypeIDs[mediaType.TypeID]; ok {
			continue
		}
		// 根据 type_pid 进行分类; 如果有大类过滤则过滤掉大类
		if _, ok := filterTypePid[strconv.Itoa(mediaType.TypePID)]; !ok {
			continue
		}
		if mediaType.TypeID < 5 {
			continue
		}
		mediaType.BaseUrl = fmt.Sprintf(baseUrl, mediaType.TypeID)
		if _, ok := data[mediaType.TypePID]; !ok {
			data[mediaType.TypePID] = make([]MediaType, 0)
		}
		data[mediaType.TypePID] = append(data[mediaType.TypePID], mediaType)
	}
	m.UrlList = data
}
