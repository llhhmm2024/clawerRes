package database

import (
	"crawler/global"
	"crawler/internal/model"
	"crawler/internal/utils"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Operator struct {
	Video       model.TblVideo
	SingleVideo []*model.TblSingleVideo
}

var MysqlChan = make(chan Operator, 1000)

func SaveToMysql(videoData model.TblVideo, singleVideo []*model.TblSingleVideo) Operator {
	return Operator{Video: videoData, SingleVideo: singleVideo}
}

func MysqlPipeline() {
	for v := range MysqlChan {
		DoMysql(v.Video, v.SingleVideo)
	}
	global.MysqlCahn <- "done"
}

func DoMysql(videoData model.TblVideo, singleVideo []*model.TblSingleVideo) {
	now := time.Now()
	videoData.CreateTime = now
	videoData.UpdateTime = now
	videoData.EpisodeUpdateTime = now

	// 更新参数
	// vals := []string{"name", "name_spell", "state", "video_type", "video_num", "screen_year", "screen_time", "pic_w",
	// 	"current_episode", "episode_status", "area_name", "play_link", "sub_title", "extend_id", "description",
	// 	"tags_name", "category_name", "director_name", "actor_name"}
	// err := db.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "id"}},
	// 	DoUpdates: clause.AssignmentColumns(vals),
	// }).Create(&videoData).Error
	// if err != nil {
	// 	fmt.Println("insert video: ", err)
	// }
	insertOrUpdateVideo(videoData)
	video := make(map[string]interface{})
	query := map[string]interface{}{"extend_id": videoData.ExtendId, "source": videoData.Source}
	rows := db.Table("tbl_video").Select("id").Where(query).Find(&video).RowsAffected
	if rows != 0 {
		var group []model.TblSingleVideo
		vid := video["id"].(uint64)
		for _, single := range singleVideo {
			single.VideoID = vid
			single.State = 1
			single.CreateTime = now
			single.UpdateTime = now
			single.GlobalVid = utils.GenerateUuid()
			group = append(group, *single)
		}
		insertOrUpdate(group)
	}
}

// 插入 video 数据
func insertOrUpdateVideo(videoData model.TblVideo) {
	var existData model.TblVideo
	err := db.Model(&model.TblVideo{}).Where("extend_id = ? AND source = ?",
		videoData.ExtendId, videoData.Source).First(&existData).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := db.Create(&videoData).Error; err != nil {
			fmt.Println("insert video err: ", err)
		}
	} else if existData.PlayLink != videoData.PlayLink || existData.State != videoData.State ||
		existData.Name != videoData.Name || existData.NameSpell != videoData.NameSpell ||
		existData.ScreenYear != videoData.ScreenYear || existData.AreaName != videoData.AreaName ||
		existData.ScreenTime != videoData.ScreenTime || existData.PicW != videoData.PicW ||
		existData.CurrentEpisode != videoData.CurrentEpisode || existData.EpisodeStatus != videoData.EpisodeStatus ||
		existData.SubTitle != videoData.SubTitle || existData.TagsName != videoData.TagsName ||
		existData.Description != videoData.Description ||
		existData.LanguageName != videoData.LanguageName || existData.CategoryName != videoData.CategoryName ||
		existData.DirectorName != videoData.DirectorName || existData.ActorName != videoData.ActorName {
		videoData.ID = existData.ID
		if err := db.Updates(&videoData).Error; err != nil {
			fmt.Println("update video err: ", err)
		}
	}
}

// 插入singleVideo数据
func insertOrUpdate(group []model.TblSingleVideo) {
	db.AutoMigrate(model.TblSingleVideo{})

	transaction := db.Begin()
	for _, item := range group {
		var existData model.TblSingleVideo
		err := transaction.Model(&model.TblSingleVideo{}).Where("video_id = ? AND source = ? AND episode = ?",
			item.VideoID, item.Source, item.Episode).First(&existData).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := transaction.Create(&item).Error; err != nil {
				transaction.Rollback()
				fmt.Println("single insert err: ", err)
			}
		} else if existData.PlayLink != item.PlayLink || existData.State != item.State ||
			existData.Episode != item.Episode || existData.Type != item.Type ||
			existData.VideoName != item.VideoName || existData.Name != item.Name ||
			existData.SourcePlaySchema != item.SourcePlaySchema {
			item.ID = existData.ID
			item.VideoID = existData.VideoID
			item.GlobalVid = existData.GlobalVid
			item.CreateTime = existData.CreateTime
			if err := transaction.Save(&item).Error; err != nil {
				transaction.Rollback()
			}
		}
	}
	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		fmt.Println("single commit err: ", err)
	}
}
