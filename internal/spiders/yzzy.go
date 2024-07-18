package spiders

import (
	"crawler/configs"
	"fmt"
	"strconv"
)

/*
电视剧：国产电视剧、短剧、韩剧、欧美
电影：除了纪录片、伦理以外其他都要
综艺：只要大陆综艺
动漫：国产、日韩动漫
*/
const (
	yyzyUrl = "https://api.1080zyku.com/inc/apijson.php?ac=detail&t=%s"
)

func (m *Media) initYzzy() {
	list := make(map[string][]string)
	list["movie"] = []string{"5", "6", "7", "8", "9", "10", "11", "41"}
	list["tv"] = []string{"14", "12", "14", "15"}
	list["zy"] = []string{"62"}
	list["dm"] = []string{"66", "67"}
	list["dj"] = []string{"83"}

	cfg := configs.Cfg.YzzyList
	filterTypePid := make(map[string]bool, len(cfg))
	for _, v := range cfg {
		filterTypePid[v] = true
	}

	// 根据 type_pid 进行分类
	sourceUrlList := make(map[int][]MediaType)
	for k, v := range list {
		switch k {
		case "movie":
			sourceUrlList[1] = sorting(v)
		case "tv":
			sourceUrlList[2] = sorting(v)
		case "dm":
			sourceUrlList[3] = sorting(v)
		case "zy":
			sourceUrlList[4] = sorting(v)
		case "dj":
			sourceUrlList[83] = sorting(v)
		}
	}
	// var data map[int][]MediaType
	data := make(map[int][]MediaType)
	for k, v := range sourceUrlList {
		if _, ok := filterTypePid[strconv.Itoa(k)]; ok {
			data[k] = v
		}
	}
	m.UrlList = data
}

func sorting(list []string) (result []MediaType) {
	for _, v := range list {
		tid, _ := strconv.Atoi(v)
		result = append(result, MediaType{TypeID: tid, BaseUrl: fmt.Sprintf(yyzyUrl, v)})
	}
	return
}
