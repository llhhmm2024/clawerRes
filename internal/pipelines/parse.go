package pipelines

import (
	"crawler/global"
	"crawler/internal/model"
	"crawler/internal/utils"
	"crawler/internal/utils/http"
	"crawler/pkg/database"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func Parse(rawUrl string) {
	data := http.HttpToJson(rawUrl)
	if data == "" {
		fmt.Println("empty data: ", rawUrl)
		return
	}
	var res model.Response
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(res.List) == 0 {
		return
	}

	for _, v := range res.List {
		ProcessVideo(v, rawUrl)
	}
}

// 组合专辑
func ProcessVideo(data model.Vod, url string) {
	vodId := utils.AssertTypeIdType(data.VodID)
	fmt.Println("video:" + data.VodName + " " + "vodId: " + vodId)
	var v model.TblVideo
	v.Name = data.VodName
	v.SubTitle = data.VodSub
	v.ScreenYear = data.VodYear
	v.AreaName = data.VodArea
	tags := utils.FormatTag(data.VodClass)
	v.TagsName = tags
	v.CategoryName = tags
	v.DirectorName = utils.FormatStr(data.VodDirector)
	v.ActorName = utils.FormatStr(data.VodActor)
	v.PicW = data.VodPic
	v.LanguageName = data.VodLang
	v.Source = processItemSource(data.VodPlayFrom)
	v.VideoType = processItemVodType(data.TypeId, v.Source)
	desc := strings.Trim(data.VodContent, "\u3000")
	vodEn := data.VodEnname
	state := 1
	if v.Source == global.Ffzy_Source {
		desc = strings.Trim(data.VodBlurb, "\u3000")
		vodEn = data.VodEnFF
		state = data.VodStatus
	}
	v.State = state
	v.NameSpell = vodEn
	v.Description = desc
	v.IsPay = -1
	v.SysAdd = 1
	// 处理单视频合集
	var singleList []*model.TblSingleVideo
	// 分割处理播放地址
	vodPlayUrl := processItemVodPlayUrl(data.VodPlayURL, data.VodPlayNote, v.Source)
	for i := 1; i <= len(vodPlayUrl); i++ {
		var s model.TblSingleVideo
		s.VideoName = data.VodName
		s.Type = v.VideoType
		s.Source = v.Source
		// 处理 playUrl
		item, err := processItemPlayUrl(strings.TrimSpace(vodPlayUrl[i-1]), v.Source)
		if err != nil {
			return
		}
		playlink := utils.IsHttps(item.Playlink)
		if playlink == "" {
			continue
		}
		s.PlayLink = playlink
		s.Name = item.Name
		if v.VideoType == 4 {
			s.Episode = item.Name
		} else {
			s.Episode = strconv.Itoa(i)
		}
		s.IsPay = -1
		singleList = append(singleList, &s)
	}
	done := make(chan bool)
	go DownloadTask(singleList, done)
	result := <-done
	total := len(singleList)
	if !result || total == 0 {
		return
	}
	v.EpisodeStatus, v.Episodes = processItemVodStatus(data.VodRemarks, total, v.VideoType)
	v.Score = data.VodScore
	v.VideoNum = total
	v.PlayLink = singleList[0].PlayLink
	v.CurrentEpisode = strconv.Itoa(total)
	v.ExtendId = vodId
	// database.DoMysql(v, singleList)
	database.MysqlChan <- database.SaveToMysql(v, singleList)
}
