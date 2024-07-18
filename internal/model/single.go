package model

import (
	"strconv"
	"time"
)

type TblSingleVideo struct {
	ID               uint64    `json:"id" gorm:"column:id"`                                 // 主键,自增
	Name             string    `json:"name" gorm:"column:name"`                             // 每集视频名称
	VideoName        string    `json:"video_name" gorm:"column:video_name"`                 // 所属视频名称
	State            int       `json:"state" gorm:"column:state"`                           // 视频状态 1-正常 -1 停用
	Type             int       `json:"type" gorm:"column:type"`                             // 视频资源类型 类型 fengxing 1-正片 2-花絮 3-预告 4-音乐或MTV 5-资讯
	Episode          string    `json:"episode" gorm:"column:episode"`                       // 集数或期数 电视剧或综艺指定当前视频是第几集或期数,电影默认是1
	VideoID          uint64    `json:"video_id" gorm:"column:video_id"`                     // 所属视频id video.id
	GlobalVid        string    `json:"global_vid" gorm:"column:global_vid"`                 // 全局视频id, 采用UUID，去掉-符号
	PlayLink         string    `json:"play_link" gorm:"column:play_link"`                   // 视频播放地址 视频地址
	DurationSecond   string    `json:"duration_second" gorm:"column:duration_second"`       // 视频时长（秒）
	IsPay            int       `json:"is_pay" gorm:"column:is_pay"`                         // 免费付费标志 -1免费 1付费
	CreateTime       time.Time `json:"create_time" gorm:"column:create_time"`               // 增加时间
	UpdateTime       time.Time `json:"update_time" gorm:"column:update_time"`               // 修改时间
	Source           string    `json:"source" gorm:"type:varchar(200)"`                     // 源站点信息
	PicW             string    `json:"pic_w" gorm:"column:pic_w"`                           // 横图
	PicH             string    `json:"pic_h" gorm:"column:pic_h"`                           // 竖图
	SourcePlaySchema string    `json:"source_play_schema" gorm:"column:source_play_schema"` // 源站播放地址
}

func (m *TblSingleVideo) TableName() string {
	return "tbl_single_video"
}

func (m *TblSingleVideo) Isfee(n string) {
	ispay, _ := strconv.Atoi(n)
	m.IsPay = ispay
}
