package spiders

import (
	"crawler/configs"
	"crawler/global"
	"crawler/internal/model"
	"crawler/internal/pipelines"
	"crawler/internal/utils/http"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type Media struct {
	Source  string
	Model   string
	UrlList map[int][]MediaType
	Page    int
}

type MediaType struct {
	TypeID   int    `json:"type_id"`
	TypePID  int    `json:"type_pid"`
	TypeName string `json:"type_name"`
	BaseUrl  string
}

func NewClient(source, model string, page int) Media {
	return Media{Source: source, Model: model, Page: page}
}

// 全量入库
func (m *Media) Run() {
	m.initUrls()
	var wg sync.WaitGroup
	// 创建带有缓冲大小的信号通道
	sem := make(chan struct{}, configs.GetMaxConcurrent())
	for typePID, types := range m.UrlList {
		fmt.Printf("start: %s ,type_pid_%d:\n", time.Now().Local().Format(time.DateTime), typePID)
		for _, mediaType := range types {
			wg.Add(1)
			go m.start(mediaType.BaseUrl, &wg, sem)
		}
	}
	wg.Wait()
	fmt.Println("finish: ", time.Now().Local().Format(time.DateTime))
}

func (m *Media) initUrls() {
	if m.Source == global.Ffzy_Source {
		m.initFeifanzy()
	} else if m.Source == global.Yzzy_Source {
		m.initYzzy()
	}
}

func (m *Media) start(baseUrl string, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	sem <- struct{}{}
	data := http.HttpToJson(baseUrl)
	if data == "" {
		fmt.Println("empty data: ", baseUrl)
		return
	}
	var res model.Response
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		fmt.Println(err)
	}
	page := res.PageCount
	switch m.Source {
	case global.Ffzy_Source:
		if m.Model != "full" {
			page = m.Page
		}
	case global.Yzzy_Source:
		if m.Model != "full" {
			page = m.Page
		}
	default:
		page = res.PageCount
	}

	for p := 1; p <= page; p++ {
		rawUrl := fmt.Sprintf("%s&pg=%d", baseUrl, p)
		fmt.Println(rawUrl)
		pipelines.Parse(rawUrl)
	}
	<-sem
}

// 指定 Video_id 入库
func (m *Media) SpecialId() {
	for _, v := range test() {
		pipelines.Parse("http://api.ffzyapi.com/api.php/provide/vod/?ac=detail&ids=68593")
		fmt.Println(v)
		// pipelines.Parse(fmt.Sprintf("https://api.1080zyku.com/inc/api_mac10.php?ac=detail&ids=%s", v))
	}
}

// url := "https://api.1080zyku.com/inc/api_mac10.php?ac=detail&ids=3275"
// pipelines.Parse(url)
// }

func test() []string {
	alist := []string{
		"56926",
	}
	// alist = []string{"55325"}
	return alist
}
