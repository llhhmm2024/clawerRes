package model

import "time"

type TblAlbum struct {
	ID           uint64    `json:"id" gorm:"column:id"`                             // 主键,自增
	Name         string    `json:"name" gorm:"column:name"`                         // 合辑名称
	NameSpell    string    `json:"name_spell" gorm:"column:name_spell"`             // 合辑名称拼音
	VideoType    int       `json:"video_type" gorm:"column:video_type"`             // 视频类型 base_video_type.id
	State        string    `json:"state" gorm:"column:state"`                       // 合辑状态 1-正常 -1停用
	Age          int16     `json:"age" gorm:"column:age"`                           // 年龄
	DirectorName string    `json:"director_name" gorm:"column:director_name"`       // 导演名称,多个name用逗号分隔
	ActorName    string    `json:"actor_name" gorm:"column:actor_name"`             // 演员/嘉宾名称 多个名称用逗号分隔
	ActorPlay    string    `json:"actor_play" gorm:"column:actor_play"`             // 主演饰演角色名称 多个角色名用逗号分隔
	DubbingName  string    `json:"dubbing_name" gorm:"column:dubbing_name"`         // 配音name,多name用逗号分隔
	CompereName  string    `json:"compere_name" gorm:"column:compere_name"`         // 主持人name,多name用逗号分隔
	PicW         string    `json:"pic_w" gorm:"column:pic_w"`                       // 竖图
	PicH         string    `json:"pic_h" gorm:"column:pic_h"`                       // 横图
	Score        string    `json:"score" gorm:"column:score"`                       // 合辑评分 满分10分
	IsPay        int       `json:"is_pay" gorm:"column:is_pay"`                     // 是否收费 -1-免费 1-收费
	ScreenYear   string    `json:"screen_year" gorm:"column:screen_year"`           // 上映年份
	ScreenTime   string    `json:"screen_time" gorm:"column:screen_time"`           // 上映时间 格式：yyyy-MM-dd
	AreaName     string    `json:"area_name" gorm:"column:area_name"`               // 地区名称 多个名称用逗号分隔
	TagsName     string    `json:"tags_name" gorm:"column:tags_name"`               // 标签名称 多个name用逗号分隔
	CategoryName string    `json:"category_name" gorm:"column:category_name"`       // 分类 多个name用逗号分隔
	Description  string    `json:"description" gorm:"column:description"`           // 描述
	VideoNum     int       `json:"video_num" gorm:"column:video_num"`               // 包含视频的数量 如一部电视剧包含4个站点的视频,则为4
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`           // 增加时间
	UpdateTime   time.Time `json:"update_time" gorm:"column:update_time"`           // 修改时间
	Ts           int       `json:"ts" gorm:"column:ts"`                             // 记录时间戳 每次操作一次都需要更新此字段内容,取当前时间戳,精度到秒
	MediaType    int       `json:"media_type" gorm:"media_type;not null;default:1"` // 合辑类型 1 长视频 2 短视频
	WritersName  string    `json:"writers_name" gorm:"writers_name;not null"`       // 编剧名称
}

func (m *TblAlbum) TableName() string {
	return "tbl_album"
}
