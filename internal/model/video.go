package model

import (
	"time"
)

type TblVideo struct {
	ID                uint64    `json:"id" gorm:"column:id"`                                         // 主键,自增
	Name              string    `json:"name" gorm:"column:name" copier:"must"`                       // 专辑名称
	NameSpell         string    `json:"name_spell" gorm:"column:name_spell" copier:"must"`           // 专辑名称拼音
	VideoType         int       `json:"video_type" gorm:"column:video_type" copier:"must"`           // 视频类型 base_video_type.id
	State             int       `json:"state" gorm:"column:state" copier:"must"`                     // 合辑状态 1-正常,-1停用
	SubTitle          string    `json:"sub_title" gorm:"column:sub_title" copier:"must"`             // 副标题
	Source            string    `json:"source" gorm:"column:source" copier:"must"`                   // 所属站点名称
	ScreenYear        string    `json:"screen_year" gorm:"column:screen_year" copier:"must"`         // 上映年份
	ScreenTime        string    `json:"screen_time" gorm:"column:screen_time" copier:"must"`         // 上映日期 格式：yyyy-MM-dd
	DirectorName      string    `json:"director_name" gorm:"column:director_name" copier:"must"`     // 导演名称,多个name用逗号分隔
	ActorName         string    `json:"actor_name" gorm:"column:actor_name" copier:"must"`           // 演员/嘉宾名称 多个名称用逗号分隔
	ActorPlay         string    `json:"actor_play" gorm:"column:actor_play"`                         // 主演饰演角色名称 多个角色名用逗号分隔
	DubbingName       string    `json:"dubbing_name" gorm:"column:dubbing_name"`                     // 配音name,多name用逗号分隔
	CompereName       string    `json:"compere_name" gorm:"column:compere_name"`                     // 主持人name,多name用逗号分隔
	AreaName          string    `json:"area_name" gorm:"column:area_name" copier:"must"`             // 地区名称 多个名称用逗号分隔
	TagsName          string    `json:"tags_name" gorm:"column:tags_name" copier:"must"`             // 标签名称 多个name用逗号分隔
	CategoryName      string    `json:"category_name" gorm:"column:category_name" copier:"must"`     // 分类 多个name用逗号分隔
	LanguageName      string    `json:"language_name" gorm:"column:language_name" copier:"must"`     // 语言
	Episodes          int       `json:"episodes" gorm:"column:episodes" copier:"must"`               // 总集数 一部电视剧的集数,综艺节目的期数
	VideoNum          int       `json:"video_num" gorm:"column:video_num"`                           // 包含单集视频的数量
	CurrentEpisode    string    `json:"current_episode" gorm:"column:current_episode" copier:"must"` // 当前集数 电视剧可能有多集电视剧,指示当前视频播到第几集
	EpisodeStatus     int       `json:"episode_status" gorm:"column:episode_status"`                 // 剧集更新状态 -1 更新中,1 已完结
	EpisodeUpdateTime time.Time `json:"episode_update_time" gorm:"column:episode_update_time"`       // 剧集更新时间
	PicW              string    `json:"pic_w" gorm:"column:pic_w" copier:"must"`                     // 横图
	PicH              string    `json:"pic_h" gorm:"column:pic_h" copier:"must"`                     // 竖图
	PlayLink          string    `json:"play_link" gorm:"column:play_link" copier:"must"`             // 视频默认播放地址
	Description       string    `json:"description" gorm:"column:description" copier:"must"`         // 视频简介
	Score             string    `json:"score" gorm:"column:score" copier:"must"`                     // 视频得分 目前采用,满分10分的分制
	ExtendId          string    `json:"extend_id" gorm:"column:extend_id" copier:"must"`             // 其他站点的 id
	IsPay             int       `json:"is_pay" gorm:"column:is_pay" copier:"must"`                   // 免费付费标志：-1 免费,0 半付费,1付费
	CreateTime        time.Time `json:"create_time" gorm:"column:create_time"`                       // 增加时间
	UpdateTime        time.Time `json:"update_time" gorm:"column:update_time"`                       // 修改时间
	SysAdd            int       `json:"sys_add" gorm:"column:sys_add"`                               // 1系统添加2 手工添加
	AlbumId           int32     `json:"album_id" gorm:"column:album_id"`                             // 合辑id
	IsFreeTime        int       `json:"is_free_time" gorm:"column:is_free_time;default:-1"`          // 是否限免 -1 否, 1 是
	WritersName       string    `json:"writers_name" gorm:"column:writers_name"`                     // 编剧名称
	UpdateCycle       string    `json:"update_cycle" gorm:"column:update_cycle"`                     // 更新周期
}

func (m *TblVideo) TableName() string {
	return "tbl_video"
}
